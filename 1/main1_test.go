package main

import "testing"

func TestString(t *testing.T) {
	data := Data{
		numDecimal:     42,
		numOctal:       052,
		numHexadecimal: 0x2A,
		pi:             3.14,
		name:           "Golang",
		isActive:       true,
		complexNum:     complex(1.0, 2.0),
	}

	expected := "var numDecimal int = 42\nvar numOctal int = 52\nvar numHexadecimal int = 2A\nvar pi float64 = 3.14\nvar name string = Golang\nvar isActive bool = true\nvar complexNum complex64 = (1+2i)"
	if result := data.String(); result != expected {
		t.Errorf("\nExpected:\n%v\n,Got:\n%v", expected, result)
	}
}

func TestToString(t *testing.T) {
	data := Data{
		numDecimal:     42,
		numOctal:       052,
		numHexadecimal: 0x2A,
		pi:             3.14,
		name:           "Golang",
		isActive:       true,
		complexNum:     complex(1.0, 2.0),
	}

	expected := "42 52 2A 3.14 Golang true (1+2i)"
	if result := data.ToString(); result != expected {
		t.Errorf("\nExpected:\n%v\n,Got:\n%v", expected, result)
	}
}

func TestHash(t *testing.T) {
	data := Data{
		numDecimal:     42,
		numOctal:       052,
		numHexadecimal: 0x2A,
		pi:             3.14,
		name:           "Golang",
		isActive:       true,
		complexNum:     complex(1.0, 2.0),
	}

	hash := data.Hash()
	if len(hash) != 64 {
		t.Errorf("Expected 64 hex hash length, but got %d", len(hash))
	}
}
