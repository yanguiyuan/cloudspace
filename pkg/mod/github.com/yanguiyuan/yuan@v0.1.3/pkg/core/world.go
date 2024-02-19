package core

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type IWorld interface {
	Run()
	Init()
	OnInit()
	OnClose()
}
type Empty struct {
}
type World struct {
	entities      map[Entity]Empty
	namedEntities map[string]Entity
}

func (w *World) AddEntity(e Entity) {
	if w.entities == nil {
		w.entities = make(map[Entity]Empty)
	}
	w.entities[e] = Empty{}
}
func (w *World) AddEntityWithName(e Entity, name string) {
	if w.entities == nil {
		w.entities = make(map[Entity]Empty)
	}
	if w.namedEntities == nil {
		w.namedEntities = make(map[string]Entity)
	}
	w.entities[e] = Empty{}
	w.namedEntities[name] = e
}
func (w *World) OnInit() {

}
func (w *World) Init() {
}
func (w *World) OnClose() {
	fmt.Println("正在关闭...")
	fmt.Println("关闭成功")
}
func (w *World) Run() {
	for e, _ := range w.entities {
		e.Startup()
		go func() {
			for {
				ticker := time.NewTicker(time.Second / 60)
				if e.Dead() {
					break
				}
				e.Update()
				select {
				case <-ticker.C:
					continue
				}
			}
		}()
	}
}

type WorldProxy struct {
	w     IWorld
	close chan os.Signal
}

func NewWorldProxy(w IWorld) *WorldProxy {
	proxy := &WorldProxy{
		w:     w,
		close: make(chan os.Signal, 1),
	}
	go func() {
		signal.Notify(proxy.close, os.Interrupt, syscall.SIGTERM)
	}()
	return proxy
}
func (wp *WorldProxy) Run() {
	wp.w.Init()
	wp.w.OnInit()
	wp.w.Run()
}
func (wp *WorldProxy) Spin() {
	wp.Run()
	select {
	case <-wp.close:
		wp.w.OnClose()
	}
}
func StartCoroutine(d time.Duration, do func()) *time.Ticker {
	ticker := time.NewTicker(d)
	go func() {
		for true {
			do()
			select {
			case <-ticker.C:
				continue
			}
		}
	}()
	return ticker
}
