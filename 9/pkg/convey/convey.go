// Package convey implements a data processing pipeline where values are passed through multiple stages:
// first they are sent through a channel, then converted to float64 and cubed, and finally collected into a slice of float64 values.
package convey

import "math"

const (
	degree = 3.0 // degree is the exponent used for cubing the float64 values
)

// FirstStep receives a slice of uint8 values and sends them through a channel
func FirstStep(slice []uint8) chan uint8 {
	ch := make(chan uint8)

	go func() {
		for _, val := range slice {
			ch <- val
		}
		close(ch)
	}()

	return ch
}

// SecondStep receives uint8 values from a channel, converts them to float64 and cubes them
func SecondStep(out <-chan uint8) chan float64 {
	ch := make(chan float64)

	go func() {
		for val := range out {
			ch <- math.Pow(float64(val), degree)
		}
		close(ch)
	}()

	return ch
}

// Output collects all the results from the float64 channel into a slice and returns it
func Output(out <-chan float64) (s []float64) {
	for cube := range out {
		s = append(s, cube)
	}

	return s
}
