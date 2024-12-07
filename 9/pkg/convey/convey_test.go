package convey

import (
	"math"
	"testing"
)

func TestAll(t *testing.T) {
	tests := []struct {
		name  string
		input []uint8
	}{
		{
			name:  "simple",
			input: []uint8{1, 2, 3, 4, 5},
		},
		{
			name:  "empty",
			input: []uint8{},
		},
		{
			name:  "single element",
			input: []uint8{3},
		},
		{
			name:  "large input",
			input: make([]uint8, 100),
		},
		{
			name:  "zero element",
			input: []uint8{0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected []float64
			for _, val := range tt.input {
				expected = append(expected, math.Pow(float64(val), degree))
			}

			got := Output(SecondStep(FirstStep(tt.input)))
			if !equal(got, expected) {
				t.Fatal("wrong cube values")
			}
		})
	}
}

func equal(first, second []float64) bool {
	if len(first) != len(second) {
		return false
	}

	for cnt, val := range first {
		if val != second[cnt] {
			return false
		}
	}

	return true
}
