// Package unique provides functionality for finding unique elements between slices.
package unique

// Find Takes two slices of strings as input.
// Returns a slice containing elements that are in the first slice but not in the second.
func Find(first, second []string) (s []string) {
	m := make(map[string]struct{})
	for _, key := range second {
		m[key] = struct{}{}
	}

	for _, key := range first {
		if _, exist := m[key]; !exist {
			s = append(s, key)
		}
	}

	return s
}
