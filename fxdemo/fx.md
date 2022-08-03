# 感觉 uber/fx 并不比 getInstance 工厂好用
不好用的点：
1. 写单元测试有点麻烦
2. 需要额外管理属性（这里的属性是：用于存储注入的依赖）
3. 不完美的按需初始化
4. 通过反射实现的注入，如果有依赖问题可能要在运行时才能被发现
5. 性能损失
5. 调试、维护略复杂(主要是调试 fx 生命周期+反射，调用链就拉得比较长的, 其它都还好)

下面就前面几点展开解释一下吧（伪代码）, 看看 V 友怎么看

## 手动维护依赖属性
fx 只能在顶层方法(app 初始化时)实现自动依赖注入(invoke 调用)，非 invoke 调用则不能直接自动注入

比如要实现调用链`server->service->db`
```
// main.go
fx.Invoke(func(s *Server){
      s.start()
})
```

为了让 service 调用依赖 db, 我们一般需要在顶层用`Provide/Module`等方法生成一份依赖关系`module.Main`，比如：

```
func NewDb() *Db{
    return initDb()
}
func NewService(db *Db) *Service{
    return &Service(db) //注入 db
}
func NewServer(s *Service) *Server{
    return &Server(s)   //注入 service
}

// 用 module 管理维护依赖关系
module.Main := fx.Module("main",
    fx.Provide(
        NewDb,
        NewService,
        NewServer,
    ),
)
```

同时，由于非 invoke 调用不可直接自动注入，所以需要手动增加属性, 用来存储所注入的依赖，比如：

```
// server.go
type struct Server{
    service *Service  //增加 service 属性，用来存储 service 依赖
}

// service.go
type struct Service{
    db *Db          // 增加 db 属性，用来存储 db 依赖
}
func (s *Service) Insert(){
    //使用 db 依赖
    s.db.Insert()
}
```
虽然 New 构造器的编写是一次性的工作，但是对依赖属性的管理，是重复性的工作：
1. 如果依赖越来越很多，我们所需要**手动**给每个对象增加的**依赖属性**就越来越多, 对象会变得越发的臃肿
    - 几乎每一层对象、每一个对象，都要加各种依赖相关的属性(除非它不用依赖)， 比如 SerivceA,SerivceB 都要添加 db 属性 **重复性工作**, **繁琐**
2. 初始化对象时，必须创建所有**依赖属性**对应的依赖对象(即使可能不会使用依赖属性)，这违反了**按需初始化**的原则(这一点比较影响单元测试效率)。
    - 比如, 因为我们的 Service 依赖 DB, 创建 Service 时就必须先创建 DB 对象, 并把 DB 依赖注入到 Service ，即使不会真正使用到 DB 对象
3. 我们只能使用 OOP, 而不能使用 Function(因为 Function 不可注入依赖)

如果使用 getInstance 就不需要手动给 Service 对象增加属性了，也不用受限在 OOP 下，而且可以做到真正的按需要初始化(不使用 DB ，就不会初始化):

    // db.go
    var _inner_db *DB
    func NewDB() *DB{
        if _inner_db == nil{
            _inner_db = connectDB()
        }
        return _inner_db
    }

    // service.go
    func (s *Service) Insert(){
        NewDB().Insert() // 直接一行流调用就可以, 且是按需要初始化的,  也可以放到普通函数中调用
    }

## 写单元测试有点麻烦
### 额外的样板代码
每一处单元测试，都要手动写这么一堆样式代码(fx.New/Module/Invoke)

```
func TestXXX(t *testing.T) {
    fx.New(
        module.Service, // 引入 modeule.Service 所有的依赖
        fx.Invoke(func(s *Service) {
            err:=s.Foo()
            // todo test ...
        }),
    )
}
```

而我更喜欢简洁的一行流

```
func TestXXX(t *testing.T) {
    err := GetInstanceService().Foo()
    // todo test
}
```
    
### 需要为单元测做额外的依赖管理
如果在单元测试的孙子、孙孙子函数里面，要调用大量的依赖, 就会比较麻烦(此场景很常见).

比如下面这个示例中，孙子函数`testGetWorkflowDef` 依赖到 Workflow 对象

    func TestWorkflow(t *testing.T) {
        fx.New(
            module.Workflow, 
            fx.Invoke(func(workflow *Workflow){
                // 创建 workflow
                wfid,err := testCreateWorkflow(workflow)
                if err!=nil{
                    t.Fatal(err)
                }
                // 完成 workflow
                testFinishWorkflow(wfid)
            }),
        )
    }

    // 创建 workflow
    func testCreateWorkflow(workflow *Workflow) (string, err){
        def, err:=testGetWorkflowDef(workflow)
        wfid, err := postCreateWorkflow(def)
        return wfid,err
    }

    // 获取 workflow 定义（孙子函数依赖 workflow ）
    func testGetWorkflowDef(workflow *Workflow) *WorkflowDef{
        def:=workflow.GenerateWorkflowDef()
        return def
    }

上面示例中，为了将`workflow`这个依赖传给孙子函数`testGetWorkflowDef`, 入口方法就将`workflow`一层一层往下传.  这样做的缺点是： 层数越多、依赖越多，就越麻烦

为了避免层层传递依赖, 我想到的，就是为单元测试也引入依赖管理:
1. 首先，可以将单元测试整个调用链`testMain->testCreateWorkflow->testGetWorkflowDef`，统一放到抽像的对象`struct WorkflowTest` 中去
2. 再借助 fx, 为单元测试对象(OOP)单独提供依赖注入

具体示例如下(避免了上例中的层层传依赖的方式):

    func TestWorkflow(t *testing.T) {
        fx.New(
            module.Workflow, 
            fx.Provide(NewWorkflowTest), // 为单测单独提供依赖
            fx.Invoke(func(workflow *Workflow, wft *WorkflowTest){
                // 创建 workflow
                wfid,err := wft.testCreateWorkflow(workflow)
                if err!=nil{
                    t.Fatal(err)
                }
                // 完成 workflow
                testFinishWorkflow(wfid)
            }),
        )
    }

    type struct WorkflowTest{
        workflow *Workflow // 增加依赖属性
    }

    // 注入workflow依赖
    func NewWorkflowTest(workflow *Workflow) *WorkflowTest{
        return &WorkflowTest{
            workflow: workflow,
        }
    }

    // 创建 workflow
    func (wft *WorkflowTest) testCreateWorkflow() (string, err){
        def, err:=wft.testGetWorkflowDef()
        wfid, err := postCreateWorkflow(def)
        return wfid,err
    }

    // 获取 workflow 定义（孙子函数依赖 workflow ）
    func (wft *WorkflowTest) testGetWorkflowDef() *WorkflowDef{
        def:=wft.workflow.GenerateWorkflowDef()
        return def
    }

可以看到，维护还是比较麻烦

> 因为每个含有单元子函数依赖的单元用例，都需要手动维护单独的依赖关系。  
> 比如,我们如果有其它的测试项，`TestTask`, `TestUser`...等等，它们都要像`TestWorkflow` 那样创建(仅仅为单测)所依赖的对象(`NewWorkflowTest`)，并且在对象中增加属性来存储依赖(像`workflow`)。

如果使用 getInstance 就不需要维护依赖关系了，也不需要去添加各种依赖相关的属性了。

### 重复维护依赖关系(可以避免，但要花点时间)
如果想单独调用一个非顶层方法，比如我想做一个针对`service.Foo`这个方法的 单元测试. 

由于 golang 不能循环依赖，所以**不能复用**入口函数的依赖定义`module.Main` (`Main->Service->Main` 就产生循环了)

只好重新维护一份依赖关系: `module.Service`, 然后在单元测试中引入：

```

// 单独维护一份依赖关系: `module.Service`
module.Service := fx.Module("service",
    fx.Provide(
        NewDb,
        NewService,
    ),
)

// 然后在单元测试中引入：`module.Service`
func TestService(t *testing.T) {
    fx.New(
        module.Service, // 引入 modeule.Service 依赖
        fx.Invoke(func(s *Service) {
            err:=s.Foo()
            if err!=nil{
                t.Fatal(err)
            }
        }),
    )
}

```

不过，如果对 fx.Module 做良好的上下分层设计也可以避免重复维护依赖关系，比如：
1. module.Main 只需要引入 module.Service, 而不必引入 module.DB
2. 因为 module.Service 已经引入了 module.DB

这需要在依赖关系的设计上，花点时间

## 小结
个人认为将 fx 用于项目中，收益比起成本，并不太划算。用于低层库则更没有必要，损失性能又增加复杂性

我想还是 getInstance 工厂更简单精暴省事：
1. 如果要变更依赖，直接修改它的实现就行了
2. 如果想支持参数，使用 map 缓存实例就行了
3. 如果需要并行再加一把 lock
