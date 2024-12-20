// This program creates channels, merges them, and collects the merged result into a slice.
package main

import (
	"fmt"

	"github.com/kossadda/go-core-task/7/pkg/chans"
)

func main() {
	channels := createChannels()

	merged, done := chans.Merge(channels...)

	slice := chans.Output(merged, done)

	fmt.Println(slice)
}

// createChannels generates a list of channels, each containing two integer elements (multiples of 1 and 2).
// It returns a slice of channels.
func createChannels() (channels []chan int) {
	for i := 1; i <= 10; i++ {
		element1 := 1 * i
		element2 := 2 * i
		channels = append(channels, chans.Create(element1, element2))
	}

	return channels
}
