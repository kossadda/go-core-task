// Package generator provides functions to generate random integers
// within a specified range and send them via a channel.
package generator

import (
	"math/rand"
	"time"
)

// Int receives a random integer from the channel.
func Int(ch chan int) int {
	return <-ch
}

// Start generates random integers within a given range and sends them to the channel.
// It will run indefinitely until a signal is received through the done channel to stop the goroutine.
func Start(min, max int, ch chan int) {
	src := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(src)

	for {
		ch <- gen.Intn(max-min+1) + min
	}
}
