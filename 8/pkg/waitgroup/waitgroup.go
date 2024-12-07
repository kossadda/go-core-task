// Package waitgroup implements a custom WaitGroup using a semaphore
// for synchronizing goroutines without using sync.WaitGroup.
package waitgroup

import (
	"sync"
)

// WaitGroup is a custom implementation of a wait group that uses a semaphore
// (a buffered channel) and a counter to wait for multiple goroutines to finish.
type WaitGroup struct {
	sem   chan struct{} // Semaphore channel to signal when all goroutines are done.
	mutex sync.Mutex    // Mutex to protect the count from concurrent access.
	count int           // Counter that keeps track of the number of active goroutines.
}

// New creates and returns a new WaitGroup instance.
func New() *WaitGroup {
	return &WaitGroup{
		sem: make(chan struct{}, 1),
	}
}

// Add modifies the count of active goroutines in the WaitGroup.
// A positive delta increases the count, and a negative delta decreases it.
func (wg *WaitGroup) Add(delta int) {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	wg.count += delta
	if wg.count < 0 {
		panic("negative WaitGroup counter")
	}
}

// Done signals that a single goroutine has finished its work and decreases the count.
func (wg *WaitGroup) Done() {
	wg.Add(-1)
	if wg.count == 0 {
		select {
		case wg.sem <- struct{}{}:
		default:
		}
	}
}

// Wait blocks until all goroutines have completed their work and the counter reaches 0.
func (wg *WaitGroup) Wait() {
	<-wg.sem
}
