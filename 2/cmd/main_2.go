// Package main is the entry point of the program.
// It demonstrates the usage of functions from the "slice" package
// for manipulating slices of integers.
package main

import (
	"fmt"

	"github.com/kossadda/go-core-task/2/pkg/slice"
)

const (
	sliceSize = 25 	// sliceSize defines the size of the generated slice.
)

func main() {
	originalSlice := slice.GenerateSlice(sliceSize)
	fmt.Println("Original slice:", originalSlice)

	evenSlice := slice.SliceExample(originalSlice)
	fmt.Println("Even slice:", evenSlice)

	updatedSlice := slice.AddElements(originalSlice, 42)
	fmt.Println("Updated slice with added element:", updatedSlice)

	copiedSlice := slice.CopySlice(originalSlice)
	fmt.Println("Copied slice:", copiedSlice)

	removedSlice := slice.RemoveElement(originalSlice, 2)
	fmt.Println("Slice after removing element:", removedSlice)
}