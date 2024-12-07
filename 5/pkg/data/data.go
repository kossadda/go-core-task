// Package data provides functionality to read slices of integers from user input.
// This package is designed to handle input validation and parsing.
package data

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// New reads a slice of integers from user input until EOF or an error occurs.
func New() (numbers []int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a slice of numbers separating them using Enter. Ctrl+D - to stop")
	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Stopped")
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, "Error while entering a number:", err)
			continue
		}

		value, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error while entering a number:", err)
			continue
		}

		numbers = append(numbers, value)
	}

	return numbers
}
