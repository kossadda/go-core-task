package waitgroup

import (
	"sync"
)

type WaitGroup struct {
	sem   chan struct{}
	mutex sync.Mutex
	count int
}

func New() *WaitGroup {
	return &WaitGroup{
		sem: make(chan struct{}, 1),
	}
}

func (wg *WaitGroup) Add(delta int) {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	wg.count += delta
	if wg.count < 0 {
		panic("negative WaitGroup counter")
	}
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
	if wg.count == 0 {
		select {
		case wg.sem <- struct{}{}:
		default:
		}
	}
}

func (wg *WaitGroup) Wait() {
	<-wg.sem
}
