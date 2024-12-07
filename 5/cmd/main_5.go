// This program accepts two slices of integers from the user, finds their intersection,
// and prints the result. If either slice is empty, the program exits with an error.
package main

import (
	"fmt"
	"os"

	"github.com/kossadda/go-core-task/5/pkg/data"
	"github.com/kossadda/go-core-task/5/pkg/intersection"
)

func main() {
	slice1 := data.New()
	if len(slice1) == 0 {
		fmt.Fprintln(os.Stderr, "First slice empty")
		os.Exit(1)
	}
	slice2 := data.New()
	if len(slice2) == 0 {
		fmt.Fprintln(os.Stderr, "Second slice empty")
		os.Exit(1)
	}

	haveUniq, uniqSlice := intersection.Find(slice1, slice2)
	fmt.Println("Bool value:", haveUniq, "\nIntersection slice:", uniqSlice)
}
