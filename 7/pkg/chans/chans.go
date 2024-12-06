package chans

import (
	"fmt"
	"sync"
)

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

func Output[T comparable](channel <-chan T, done <-chan struct{}) {
loop:
	for {
		select {
		case num := <-channel:
			fmt.Println(num)
		case <-done:
			break loop
		}
	}
}
