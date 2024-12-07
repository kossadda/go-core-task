// Package chans provides utility functions to work with channels,
// including creating channels, merging multiple channels into one, and collecting
package chans

import (
	"sync"
)

// Create creates a channel of type T, sends the provided elements into it,
// and closes the channel after all elements are sent.
func Create[T comparable](elements ...T) chan T {
	ch := make(chan T)

	go func() {
		for _, i := range elements {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

// Merge merges multiple channels of type T into a single output channel.
// It also returns a 'done' channel that signals when all input channels are processed.
func Merge[T comparable](channels ...chan T) (<-chan T, <-chan struct{}) {
	out := make(chan T)
	done := make(chan struct{})
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan T) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	return out, done
}

// Output collects values from the merged channel and returns them as a slice.
// It waits for the 'done' signal to ensure all channels are processed.
func Output[T comparable](channel <-chan T, done <-chan struct{}) []T {
	var slice []T

loop:
	for {
		select {
		case value := <-channel:
			slice = append(slice, value)
		case <-done:
			break loop
		}
	}

	return slice
}
