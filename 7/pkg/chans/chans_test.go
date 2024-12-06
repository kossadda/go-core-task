package chans

import "testing"

func TestAll(t *testing.T) {
	tests := []struct {
		name     string
		expected [][]int
	}{
		{
			name: "simple",
			expected: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
			},
		},
		{
			name: "empty channels",
			expected: [][]int{
				{},
				{},
				{},
			},
		},
		{
			name: "single element channels",
			expected: [][]int{
				{1},
				{2},
				{3},
			},
		},
		{
			name: "mixed sizes",
			expected: [][]int{
				{1, 2, 3},
				{},
				{4, 5},
				{6},
			},
		},
		{
			name: "large input",
			expected: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var channels []chan int
			m := make(map[int]struct{})

			for _, s := range tt.expected {
				channels = append(channels, Create(s...))
				for _, val := range s {
					m[val] = struct{}{}
				}
			}

			merged, done := Merge(channels...)
			slice := Output(merged, done)

			for _, ss := range slice {
				if _, ok := m[ss]; !ok {
					t.Errorf("not all values get in output")
				}
			}
		})
	}
}
