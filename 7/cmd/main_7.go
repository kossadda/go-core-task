package main

import (
	"fmt"

	"github.com/kossadda/go-core-task/7/pkg/chans"
)

func main() {
	var channels []chan int
	done := make(chan struct{})

	for i := 1; i <= 10; i++ {
		element1 := 1 * i
		element2 := 2 * i
		channels = append(channels, chans.Create(element1, element2))
	}

	merged := chans.Merge(done, channels...)

outerloop:
	for {
		select {
		case num := <-merged:
			fmt.Println(num)
		case <-done:
			break outerloop
		}
	}
}
