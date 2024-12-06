package main

import (
	"github.com/kossadda/go-core-task/7/pkg/chans"
)

func main() {
	channels := createChannels()

	merged, done := chans.Merge(channels...)

	chans.Output(merged, done)
}

func createChannels() (channels []chan int) {
	for i := 1; i <= 10000; i++ {
		element1 := 1 * i
		element2 := 2 * i
		channels = append(channels, chans.Create(element1, element2))
	}

	return channels
}
