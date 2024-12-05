// Package data provides functionality for reading and validating user input.
package data

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// New reads a slice of strings from standard input.
// Returns the slice or an error if input is invalid or reading fails.
func New() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter slice separated by space: ")
	input, err := reader.ReadString('\n')
	if err == io.EOF {
		return nil, errors.New("stopped reading from stdin")
	} else if err != nil {
		return nil, err
	}

	sl := strings.Fields(input)
	if len(sl) == 0 {
		return nil, errors.New("empty slice input")
	}

	return sl, nil
}
