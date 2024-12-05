package slice

import (
	"testing"
)

func TestGenerateSlice(t *testing.T) {
	s := GenerateSlice(10)
	if len(s) != 10 {
		t.Errorf("expected length 10, got %d", len(s))
	}
}

func TestSliceExample(t *testing.T) {
	input := Slice{1, 2, 3, 4, 5}
	expected := Slice{2, 4}
	result := SliceExample(input)

	if !equalSlices(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	input := Slice{1, 2, 3}
	expected := Slice{1, 2, 3, 42}
	result := AddElements(input, 42)

	if !equalSlices(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestCopySlice(t *testing.T) {
	input := Slice{1, 2, 3}
	result := CopySlice(input)

	if !equalSlices(result, input) {
		t.Errorf("expected %v, got %v", input, result)
	}

	result[0] = 10
	if input[0] == result[0] {
		t.Errorf("expected different values for input and result after modification")
	}
}

func TestRemoveElement(t *testing.T) {
	input := Slice{1, 2, 3, 4, 5}
	expected := Slice{1, 2, 4, 5}

	result := RemoveElement(input, 2)
	if !equalSlices(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for out-of-range index, but no panic occurred")
		}
	}()
	RemoveElement(input, 10)
}

func equalSlices(a, b Slice) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
