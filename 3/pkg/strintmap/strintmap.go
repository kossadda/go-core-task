// Package strintmap provides an implementation of a data structure for storing "string-integer" pairs.
// It supports operations like adding, removing, copying elements, checking key existence, and retrieving values by key.
package strintmap

// StringIntMap represents a data structure for storing "key-value" pairs, where the key is a string
// and the value is an integer.
type StringIntMap struct {
	sim map[string]int // underlying map storing key-value pairs
}

// New creates and returns a new instance of StringIntMap.
func New() *StringIntMap {
	return &StringIntMap{
		sim: map[string]int{},
	}
}

// Add inserts a new "key-value" pair into the map.
// If the key already exists, its value is updated.
func (s *StringIntMap) Add(key string, value int) {
	s.sim[key] = value
}

// Remove deletes the "key-value" pair with the given key from the map.
// If the key does not exist, the map remains unchanged.
func (s *StringIntMap) Remove(key string) {
	delete(s.sim, key)
}

// Copy creates and returns a new map containing all key-value pairs from the original map.
func (s *StringIntMap) Copy() map[string]int {
	m := make(map[string]int)

	for key, value := range s.sim {
		m[key] = value
	}

	return m
}

// Exists checks whether the map contains a key.
// Returns true if the key exists, false otherwise.
func (s *StringIntMap) Exists(key string) bool {
	_, exist := s.sim[key]
	return exist
}

// Get retrieves the value associated with the given key.
// It returns the value and a boolean indicating whether the key was found.
func (s *StringIntMap) Get(key string) (int, bool) {
	val, exist := s.sim[key]
	return val, exist
}
