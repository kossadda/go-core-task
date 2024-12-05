package unique

import (
	"testing"
)

func TestScan(t *testing.T) {
	tests := []struct {
		name     string
		first    []string
		second   []string
		expected []string
	}{
		{
			name:     "Common elements",
			first:    []string{"apple", "banana", "cherry"},
			second:   []string{"banana", "cherry"},
			expected: []string{"apple"},
		},
		{
			name:     "No common elements",
			first:    []string{"apple", "banana"},
			second:   []string{"fig", "grape"},
			expected: []string{"apple", "banana"},
		},
		{
			name:     "Empty first slice",
			first:    []string{},
			second:   []string{"banana", "cherry"},
			expected: []string{},
		},
		{
			name:     "Empty second slice",
			first:    []string{"apple", "banana"},
			second:   []string{},
			expected: []string{"apple", "banana"},
		},
		{
			name:     "Both slices empty",
			first:    []string{},
			second:   []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Find(tt.first, tt.second)
			if !equal(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
