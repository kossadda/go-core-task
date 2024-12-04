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

	expected := "numDecimal 42 - int\nnumOctal 52 - int\nnumHexadecimal 2A - int\npi 3.14 - float64\nname Golang - string\nisActive true - bool\ncomplexNum (1+2i) - complex64"
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
