// Package main demonstrates the usage of the `strintmap` package by implementing all the supported methods
// of the `StringIntMap` data structure and printing the results to the console.
//
// This package shows how to create a map, add key-value pairs, check for key existence,
// retrieve values by key, remove elements, copy the map, and ensure the copied map is independent of the original.
package main

import (
	"fmt"

	sim "github.com/kossadda/go-core-task/3/pkg/strintmap"
)

func main() {
	m := sim.New()

	m.Add("342", 1234)
	m.Add("567", 5678)
	m.Add("789", 91011)

	value, exists := m.Get("342")
	fmt.Printf("Get(\"342\") = %d, exists: %v\n", value, exists)

	value, exists = m.Get("999")
	fmt.Printf("Get(\"999\") = %d, exists: %v\n", value, exists)

	exists = m.Exists("567")
	fmt.Printf("Exists(\"567\") = %v\n", exists)

	exists = m.Exists("1000")
	fmt.Printf("Exists(\"1000\") = %v\n", exists)

	m.Remove("567")
	fmt.Println("After Remove(\"567\"): ")

	exists = m.Exists("567")
	fmt.Printf("Exists(\"567\") = %v\n", exists)

	copiedMap := m.Copy()
	fmt.Println("Copied map:")
	for key, value := range copiedMap {
		fmt.Printf("%s: %d\n", key, value)
	}

	m.Add("101", 2020)
	fmt.Println("After Add(\"101\", 2020):")
	fmt.Println("Original map:")
	for key, value := range m.Copy() {
		fmt.Printf("%s: %d\n", key, value)
	}
	fmt.Println("Copied map after original map modification:")
	for key, value := range copiedMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
