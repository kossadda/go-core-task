package intersection

import (
	"github.com/kossadda/go-core-task/5/pkg/util"
	"testing"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name     string
		first    []int
		second   []int
		expected []int
		expBool  bool
	}{
		{
			name:     "all",
			first:    []int{1, 2, 3},
			second:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
			expBool:  true,
		},
		{
			name:     "one intersect",
			first:    []int{2},
			second:   []int{1, 2, 3},
			expected: []int{2},
			expBool:  true,
		},
		{
			name:     "no intersect",
			first:    []int{1, 2, 3},
			second:   []int{4, 5, 6},
			expected: []int{},
			expBool:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, gotSlice := Find(tt.first, tt.second)
			if gotBool != tt.expBool {
				t.Errorf("Expected %v, got %v", tt.expBool, gotBool)
			}
			if !util.Equal(tt.expected, gotSlice) {
				t.Errorf("Expected %v, got %v", tt.expected, gotSlice)
			}
		})
	}
}
