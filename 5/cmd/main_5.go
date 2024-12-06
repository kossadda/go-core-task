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
