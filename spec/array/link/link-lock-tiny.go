type List struct {
  head, tail *MutexNode
  size int64
}

type MutexNode struct {
  Locker *sync.Mutex
  Value interface{}
  Next *MutexNode
}

func NewMutexNode(v interface{}) *MutexNode {
  n := &MutexNode{Value: v}
  n.Locker = &sync.Mutex{}
  n.Next = nil
  return n
}

func NewList() *List {
  l := &List{}
  l.head = NewMutexNode(nil)
  l.tail = l.head

  return l
}
//下面我们来看Enq，功能与方案一一致，只是在处理锁的地方有所不同，因为tail节点总是在list末尾的元素，符合我们从head开始的加锁顺序，又因为l.tail的位置始终是确定的，所以可以省略锁住前项节点的步骤；然而l.tail会在我们等待锁的时间段里被更新，所以我们需要处理l.tail被更新的情况：

func (l *List) Enq(v interface{}) bool {
    if v == nil {
        return false
    }

    tail := l.tail
    tail.Locker.Lock()
    for tail.Next != nil {
        next := tail.Next
        next.Locker.Lock()
        tail.Locker.Unlock()
        tail = next
    }

    node := NewMutexNode(v)
    tail.Next = node
    l.tail = node
    l.size++
    tail.Locker.Unlock()
    
    return true
}

//如果tail的next在取得锁时不为nil，说明tail被更新，在tail被更新之后我们需要找到当前的末尾节点，这时不能直接使用l.tail，有两点原因，一是因为这时的l.tail可能也已经被更新，二是在临时变量tail可能是非前驱节点时给l.tail加锁不能保证其一致性，而且如此一来会破坏加锁的顺序，会造成意想不到的问题。所以我们遵循加锁的顺序原则不断后推，直到找到真正的末尾节点。由此可见方案二的Enq操作最坏情况下是O(n)最好情况下是O(1)，而只要仔细想一想，在并发压力较大时这个操作几乎总是O(n)的时间开销（不过实际情况是方案二花费的时间与方案一差不多，原因在于方案一要锁住整个list开销实在太大了）

//Insert的功能也与方案一一样，因为不是锁住整个list，所以光判断size是无意义的，需要处理list被中途修改的情况；而且因为是从head开始加锁，然后锁住节点1再解锁head，以此类推，所以不会有竞争，但同样存在remove和insert一起使用时insert会失败的情况，需要注意其返回值：

func (l *List) Insert(index int64, v interface{}) bool {
    if index < 0 || v == nil {
        return false
    }

    current := l.head
    current.Locker.Lock()
    for i := int64(0); i <= index-1; i++ {
        next := current.Next
    if next == nil {
      // 如果index前的某个节点为nil，那么说明链表可能被修改了，没有index个节点，insert失败
      if index < index - 1 {
        current.Locker.Unlock()
        return false
      }

      break
    }
    next.Locker.Lock()
    current.Locker.Unlock()
    current = next
    }

    node := NewMutexNode(v)
    node.Next = current.Next
    current.Next = node
    l.size++
    current.Locker.Unlock()

    return true
}
