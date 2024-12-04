// Package main implements a program to create variables of different data
// types, convert them to a string, hash them, and test the functionality.

package main

import (
	"fmt"

	"github.com/kossadda/go-core-task/1/pkg/types"
)

func main() {
	data := types.Data{
		NumDecimal:     42,
		NumOctal:       052,
		NumHexadecimal: 0x2A,
		Pi:             3.14,
		Name:           "Golang",
		IsActive:       true,
		ComplexNum:     complex(1.0, 2.0),
	}

	fmt.Println(&data)
	fmt.Println("Sha256 -", data.Hash())
}
