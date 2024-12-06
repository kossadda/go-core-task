// Package util provides utility functions for comparing slices.
// This package is intended to be reusable across various components of the program.
package util

// Equal checks whether two slices of comparable type are equal in length and content.
func Equal[T comparable](a, b []T) bool {
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
