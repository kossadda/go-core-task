package waitgroup

import (
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	wg := New()

	tasks := 3
	wg.Add(tasks)

	for i := 0; i < tasks; i++ {
		go func(taskID int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
		}(i)
	}

	wg.Wait()
	if wg.count != 0 {
		t.Errorf("Expected counter to be 0, got %d", wg.count)
	}
}

func TestNegativeCounterPanic(t *testing.T) {
	wg := New()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on negative counter, but no panic occurred")
		}
	}()
	wg.Done()
}

func TestWaitWithoutAdd(t *testing.T) {
	wg := New()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		t.Error("Wait() returned without any tasks being added")
	case <-time.After(time.Millisecond * 100):
	}
}
