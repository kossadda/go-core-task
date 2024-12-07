package generator

import (
	"testing"
)

const (
	maxValue = 1000
	minValue = -1000
)

func TestInt(t *testing.T) {
	ch := make(chan int)
	defer close(ch)
	done := make(chan struct{})
	defer close(done)

	go Start(minValue, maxValue, ch, done)

	value := Int(ch)

	if value < minValue || value > maxValue {
		t.Errorf("Expected a value between %d and %d, but got %d", minValue, maxValue, value)
	}
}

func TestStart(t *testing.T) {
	ch := make(chan int)
	defer close(ch)
	done := make(chan struct{})
	defer close(done)

	go Start(minValue, maxValue, ch, done)

	for i := 0; i < 10; i++ {
		value := <-ch
		if value < minValue || value > maxValue {
			t.Errorf("Generated value %d out of range [%d, %d]", value, minValue, maxValue)
		}
	}
}

func TestEdgeCase(t *testing.T) {
	ch := make(chan int)
	defer close(ch)
	done := make(chan struct{})
	defer close(done)

	go Start(minValue, maxValue, ch, done)

	for i := 0; i < 10; i++ {
		value := <-ch
		if value == minValue || value == maxValue {
			t.Logf("Edge case: Generated value %d within expected range", value)
		}
	}
}
