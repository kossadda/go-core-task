// Package intersection provides functionality to find common elements between two slices of integers.
// This package implements efficient algorithms using maps for quick lookup.
package intersection

// Find computes the intersection of two slices and returns a boolean indicating
// if there are common elements and the intersection slice.
func Find(first, second []int) (bool, []int) {
	var s []int
	m := make(map[int]struct{})
	for _, key := range second {
		m[key] = struct{}{}
	}

	for _, key := range first {
		if _, exist := m[key]; exist {
			s = append(s, key)
		}
	}

	return len(s) > 0, s
}
