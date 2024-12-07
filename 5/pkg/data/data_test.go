package data

import (
	"os"
	"testing"

	"github.com/kossadda/go-core-task/5/pkg/util"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{
			name:     "first",
			input:    "123\n342\n543\n4\n",
			expected: []int{123, 342, 543, 4},
		},
		{
			name:     "second",
			input:    "13\n3\n",
			expected: []int{13, 3},
		},
		{
			name:     "empty",
			input:    "",
			expected: []int{},
		},
		{
			name:     "incorrect",
			input:    "16\n13f\n",
			expected: []int{16},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSlice := oneCase(t, tt.input)
			if !util.Equal(gotSlice, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, gotSlice)
			}
		})
	}
}

func oneCase(t *testing.T, input string) []int {
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	_, err = w.Write([]byte(input))
	if err != nil {
		t.Fatalf("Failed to write to pipe: %v", err)
	}
	w.Close()

	os.Stdin = r

	return New()
}
