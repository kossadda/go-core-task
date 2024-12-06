package chans

import "fmt"

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

func Merge[T comparable](done chan<- struct{}, channels ...chan T) <-chan T {
	out := make(chan T)
	rangeRoutines := make(chan struct{})

	for _, ch := range channels {
		go func(c <-chan T) {
			for v := range c {
				out <- v
			}
			rangeRoutines <- struct{}{}
		}(ch)
	}

	go func() {
		for i := 0; i < len(channels); i++ {
			<-rangeRoutines
		}
		close(rangeRoutines)
		done <- struct{}{}
		fmt.Println("Closed")
	}()

	return out
}
