// Package main is the entry point of the application.
// It reads two slices of strings from the user and finds elements in the first slice
// that are not present in the second slice.
package main

import (
	"fmt"
	"os"

	"github.com/kossadda/go-core-task/4/pkg/data"
	"github.com/kossadda/go-core-task/4/pkg/unique"
)

func main() {
	slice1, err := data.New()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "data.New: %v\n", err)
		return
	}

	slice2, err := data.New()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "data.New: %v\n", err)
		return
	}

	fmt.Println("Unique elements:", unique.Find(slice1, slice2))
}
