package waitgroup

import (
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	wg := New()

	tasks := 3
	wg.Add(int64(tasks))

	for i := 0; i < tasks; i++ {
		go func(taskID int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 100)
		}(i)
	}

	wg.Wait()
	if !wg.state.EqZero() {
		t.Errorf("Expected counter to be 0, got %d", wg.state.cnt)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on negative counter, but no panic occurred")
		}
	}()
	wg.Done()
}
