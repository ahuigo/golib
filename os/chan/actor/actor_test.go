package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"sync"
	"testing"
	"time"
)

// 关于中介者模式，请参考：http://timd.cn/design-pattern/mediator/；
// 关于模板方法模式，请参考：http://timd.cn/design-pattern/template-method/
/**
一个 Actor 是一个并发执行的实体，它可以接收消息、处理消息、发送消息，以及启动和停止。它分成
1. state：Actor 的状态
2. behavior：Actor 的行为, 改变状态的方法
3. message：Actor 的消息
	1. 发送: 发送消息给其他 Actor（可指定Actor或随机）
	2. 接收


*/
type Registry struct {
	actors map[string][]Actor
}

func (r *Registry) Start() {
	for _, actors := range r.actors {
		for _, actor := range actors {
			go actor.Start()
		}
	}
}

func (r *Registry) AddActor(a Actor) {
	if _, found := r.actors[a.GetName()]; found {
		r.actors[a.GetName()] = append(r.actors[a.GetName()], a)
	} else {
		r.actors[a.GetName()] = []Actor{a}
	}
}

func (r *Registry) GetActors(name string) []Actor {
	if actors, found := r.actors[name]; found {
		return actors
	}
	return nil
}

func (r *Registry) SendMessage(to string, index int, message MessageType, waitFor time.Duration) bool {
	if actors, found := r.actors[to]; found {
		if index < 0 {
			index = rand.Int() % len(actors)
		}
		return actors[index].Receive(message, waitFor)
	} else {
		return false
	}
}

func (r *Registry) Stop() {
	for _, actors := range r.actors {
		for _, actor := range actors {
			actor.Stop()
		}
	}
}

func NewRegistry() *Registry {
	return &Registry{actors: map[string][]Actor{}}
}

// MessageType 表示消息类型
type MessageType = any

// Actor 表示 Actor
type Actor interface {
	// Start 启动 Actor
	Start()
	// Receive 接收消息
	Receive(MessageType, time.Duration) bool
	// Stop 停止 Actor
	Stop()
	// GetName 获取 Actor 的名字
	GetName() string
}

// BaseActor 是所有 Actor 的基类
type BaseActor struct {
	mailBox            chan MessageType
	stopCh             chan struct{}
	closeCallback      func()
	processMessageFunc func(MessageType, *Registry)
	getNameFunc        func() string
	registry           *Registry
}

func (b *BaseActor) Start() {
	log.Printf("%s started", b.GetName())
	defer b.closeCallback()
	for {
		select {
		case message, ok := <-b.mailBox:
			if !ok {
				goto END
			}
			b.processMessageFunc(message, b.registry)
		default:
			select {
			case message, ok := <-b.mailBox:
				if !ok {
					goto END
				}
				b.processMessageFunc(message, b.registry)
			case <-b.stopCh:
				goto END
			}
		}
	}

END:
	log.Printf("%s stopped\n", b.GetName())
}

func (b *BaseActor) Receive(message MessageType, waitFor time.Duration) bool {
	select {
	case b.mailBox <- message:
		return true
	default:
		select {
		case b.mailBox <- message:
			return true
		case <-time.After(waitFor):
			return false
		}
	}
}

func (b *BaseActor) Stop() {
	close(b.stopCh)
}

func (b *BaseActor) GetName() string {
	return b.getNameFunc()
}

func NewBaseActor(mailBoxSize int, closeCallback func(),
	processMessageFunc func(MessageType, *Registry), getNameFunc func() string,
	registry *Registry) *BaseActor {
	b := &BaseActor{
		mailBox:            make(chan MessageType, mailBoxSize),
		stopCh:             make(chan struct{}),
		closeCallback:      closeCallback,
		processMessageFunc: processMessageFunc,
		getNameFunc:        getNameFunc,
		registry:           registry,
	}
	registry.AddActor(b)
	return b
}

type Assembler struct {
	*BaseActor
	no int
}

func (a *Assembler) processMessage(message MessageType, registry *Registry) {
	log.Printf("%s[%d] is processing message: %s", a.GetName(), a.no, message)
	sent := registry.SendMessage("sinker", -1, message, time.Second)
	log.Printf("%s[%d] sent message[%s] to sinker: %v", a.GetName(), a.no, message, sent)
}

func NewAssembler(mailBoxSize int, closeCallback func(), registry *Registry, no int) *Assembler {
	a := &Assembler{no: no}
	a.BaseActor = NewBaseActor(mailBoxSize, closeCallback, a.processMessage, func() string {
		return "assembler"
	}, registry)
	return a
}

type Sinker struct {
	*BaseActor
	no int
}

func (s *Sinker) processMessage(message MessageType, _ *Registry) {
	log.Printf("%s[%d] is processing message: %s", s.GetName(), s.no, message)
}

func NewSinker(mailBoxSize int, closeCallback func(), registry *Registry, no int) *Sinker {
	s := &Sinker{no: no}
	s.BaseActor = NewBaseActor(mailBoxSize, closeCallback, s.processMessage, func() string {
		return "sinker"
	}, registry)
	return s
}

func TestActor(t *testing.T) {
	var wg sync.WaitGroup
	registry := NewRegistry()

	var assemblers []*Assembler
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		assemblers = append(assemblers, NewAssembler(100, wg.Done, registry, i))
	}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		NewSinker(100, wg.Done, registry, i)
	}

	registry.Start()

	for i := 1; i <= 100; i++ {
		assemblers[i%len(assemblers)].Receive(fmt.Sprintf("message-%d", i), time.Second)
	}

	time.Sleep(time.Second)
	registry.Stop()
	wg.Wait()
}
