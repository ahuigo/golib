# 时间轮
- 本目录是一个时间轮的实现
- 另外一个更好的实现是：git@github.com:ahuigo/go-timewheel.git
    - Refer: 完全兼容golang定时器的高性能时间轮实现(go-timewheel) https://xiaorui.cc/archives/6160

## 定时器(时间轮: set popCh)

### curd:

    // init timeWheel
    tw=timer.NewTimer()
        // i<256
        t.t1[i] = newTimeHead(1, i)
        // i<4,j<64
        t.t2Tot5[i][j] = newTimeHead(i+2, j)
        // curTimePoint
        t.curTimePoint = get10Ms()

    // add task
    timer := tw.ScheduleFunc(time.Second*time.Duration(userExpire=api.CheckInterval), func() {
        println(api)
    })
        1. node.expire = e10ms(userExpire)+jiffies
            userExpire 间隔时间，比如1s
            e10ms(userExpire) 就是1s/10ms = 100
        2. tw.add(node, jiffies=t.jiffies)
            1. idx=expire-jiffies, level=1, index =0
                // jiffies 是当前时间片计数，
                // expire 是下次执行的时间片
                // idx 是离下次执行的时间片差值(expire-jiffies)
            2. if idx < nearSize: //2^8
                index = expire & nearSize
                head = t.t1[index]
                // 255个时间片
            3. else:
                1. max = maxVal() // 1<<32-1
                2. for i in range(3):
                    2. if idx>max:  // 防溢出, 最大的seq
                        idx=max
                        expire = idx+jiffies
                    3. idx < levelMax(i+1):
                        index = expire >> (nearShift+i*levelShift) & levelMask
                            // expire >> (8+i*6) & 63
                            // level=0: 每index加1，时间片增 2^8 = 256
                            // level=1: 每index加1，时间片增 2^(8+6) = 1634
                            // level=3: 每index加1，时间片增 2^(8+6*3) = 0.67亿
                            // index小于256, 最大时间片约170亿(5年)
                        head = t.t2Tot5[i][index]
                        level = i + 2
                        break;
                3. head.lockPushBack(node, level, index)
                    // 把node 加入head 链表(list)
                    // 如果已经stop停止，就不加
                    // node.version = head.version
                    // node.list = head


    // delete task
   //timer.Stop() 


    // run 
    go tw.Run()

### newTimeHead(level, index)

    // version:
    // level<<(32+16)|index<<32|
    // |---16bit---|---16bit---|------32bit-----|
    // |---level---|---index---|-------seq------|
    // seq: jiffies: 32bit
    //     --6bit--6bit--6bit--6bit--8bit----
    /*     -- 4*levelShift -------nearShift--- 
    level1: 
        t1[index] 是最近的盘子, nearSize是256个
        t.t2Tot5[i][j] 有4个盘子(4个level: 2-6, levelSize: 64)
    const:
        nearMask: 1<<8 -1
        levelMask: 1<<6 -1 = 63
    分级：
        // level 在near盘子里就是1, 在T2ToTt[0]盘子里就是2起步
        // index 就是各自盘子的索引值
    maxVal of seq:
        1<<32 -1
        1<<(8+4*6) -1
        1 << uint32(nearShift+4*levelShift)) - 1
    levelMax(index):
        1 << (8+index*6)
        1 << uint32(nearShift+index*levelShift)
    */

### Time:TimeNode
node.version: 

    1. 每次cascade 移动，都会加1
    1. 每次moveAndExec 移动，都会加1
    2. 新加入的node都要设置为head list的version：
        	atomic.StoreUint64(&node.version, atomic.LoadUint64(&t.version))
    3. version作用：判断list中

### tw.Run: timeWheel.run
每隔一个时间片，就调用一下tw.run

    1. t=tw=timer.NewTimer(),
    2. moveAndExec()
        1. index := t.jiffies & nearMask
        // 小盘子256个槽位被256时间片消耗完了后，就从大盘子里面移动一些节点过来
        2. if index = 0:   
            1. for i in range(3):
                // 将大盘level，移动到小盘level
                // 先移动小盘level, index=rang(0,64)
                // 当小盘level的index!=0时，说明小盘还没有消耗完时间片，大盘不用移动（break）
                index2 := t.index(i)
                t.cascade(i, int(index2))
                if index2!=0:
                    break
        3. t.jiffies+=1
        1. head := newTimeHead(0, 0)
        2. t1 := t.t1[index]
        3. t1.ReplaceInit(&head.Head)
        4. head.ForEach:
            1. val = list(offset)
            2. val.callback()
            2. t.add(val, jiffies)
