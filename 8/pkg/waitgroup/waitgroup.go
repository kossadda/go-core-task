package waitgroup

import (
	"runtime"
	"sync/atomic"
)

func New() *WaitGroup {
	return &WaitGroup{}
}

type WaitGroup struct {
	_     noCopy
	state counter
}

type noCopy struct{}

type counter struct {
	_   noCopy
	cnt int64
}

func (c *counter) Add(delta int64) {
	atomic.AddInt64(&c.cnt, delta)
}

func (c *counter) Sub(delta int64) {
	atomic.AddInt64(&c.cnt, -delta)
}

func (c *counter) EqZero() bool {
	return c.cnt == 0
}

func (wg *WaitGroup) Add(delta int64) {
	wg.state.Add(delta)
}

func (wg *WaitGroup) Done() {
	if atomic.LoadInt64(&wg.state.cnt) <= 0 {
		panic("negative WaitGroup counter")
	}
	wg.state.Sub(1)
}

func (wg *WaitGroup) Wait() {
	for !wg.state.EqZero() {
		runtime.Gosched()
	}
}
