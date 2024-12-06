// This program generates and prints 10 random integers within a specified range
// using the generator package. It creates a goroutine that generates random numbers
// and sends them through a channel, which are then retrieved and printed in the main function.
package main

import (
	"fmt"
	"github.com/kossadda/go-core-task/6/pkg/generator"
)

const (
	maxValue = 1000  // Maximum random value.
	minValue = -1000 // Minimum random value.
)

func main() {
	ch := make(chan int)
	defer close(ch)
	done := make(chan struct{})
	defer close(done)

	go generator.Start(minValue, maxValue, ch, done)

	for i := 0; i < 10; i++ {
		fmt.Println(generator.Int(ch))
	}
}
