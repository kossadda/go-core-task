// Package slice provides utility functions for working with slices of integers.
// It includes functionalities for generating slices, filtering even numbers,
// appending elements, copying slices, and removing elements by index.
package slice

import (
	"math/rand"
	"time"
)

type Slice []int

const (
	maxValue         = 1000    // Maximum random value.
	minValue         = -1000   // Minimum random value.
	conditionCapSize = 1 << 10 // Capacity threshold for doubling.
	conditionCapExt  = 1.25    // Capacity multiplier for large slices.
)

// GenerateSlice creates a slice of random integers of size n.
func GenerateSlice(n int) Slice {
	src := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(src)

	slice := make(Slice, n)

	for i := 0; i < n; i++ {
		slice[i] = randGen.Intn(maxValue-minValue) + minValue
	}

	return slice
}

// SliceExample returns a new slice containing only even numbers from the input.
func SliceExample(slice Slice) Slice {
	var evenSlice Slice

	for _, num := range slice {
		if num%2 == 0 {
			evenSlice = append(evenSlice, num)
		}
	}

	return evenSlice
}

// AddElements appends a value to the slice and returns the updated slice.
func AddElements(slice Slice, val int) (s Slice) {
	l := len(slice)
	newCap := func() int {
		if len(slice) < conditionCapSize {
			return cap(slice) * 2
		} else {
			return int(float64(cap(slice)) * conditionCapExt)
		}
	}

	if l+1 < cap(slice) {
		s = slice[:l+1]
	} else {
		s = make(Slice, l+1, newCap())
	}

	copy(s, slice)
	s[l] = val

	return s
}

// CopySlice creates and returns a copy of the input slice.
func CopySlice(slice Slice) (s Slice) {
	s = make(Slice, len(slice), cap(slice))

	for i := range slice {
		s[i] = slice[i]
	}

	return s
}

// RemoveElement removes an element at the given index and returns the updated slice.
// Panics if the index is out of bounds.
func RemoveElement(slice Slice, index uint) (s Slice) {
	if int(index) >= len(slice) {
		panic("removeElement: out of range")
	}

	return append(slice[:index], slice[index+1:]...)
}
