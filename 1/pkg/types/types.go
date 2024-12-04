package types

import (
	"crypto/sha256"
	"fmt"
)

const (
	saltExample = "go-2024" // salt for hash
)

// Data structure to store different types of data variables.
type Data struct {
	NumDecimal     int       // Decimal number
	NumOctal       int       // Octal number
	NumHexadecimal int       // Hexadecimal number
	Pi             float64   // Floating-point number
	Name           string    // String
	IsActive       bool      // Boolean value
	ComplexNum     complex64 // Complex number
}

// String method implements the Stringer interface for the Data structure.
// It returns a string representation of all variables in the structure,
// including both their types and values.
func (d *Data) String() string {
	return fmt.Sprintf(
		"var numDecimal %T = %d\n"+
			"var numOctal %T = %o\n"+
			"var numHexadecimal %T = %X\n"+
			"var pi %T = %g\n"+
			"var name %T = %s\n"+
			"var isActive %T = %t\n"+
			"var complexNum %T = %v",
		d.NumDecimal, d.NumDecimal,
		d.NumOctal, d.NumOctal,
		d.NumHexadecimal, d.NumHexadecimal,
		d.Pi, d.Pi,
		d.Name, d.Name,
		d.IsActive, d.IsActive,
		d.ComplexNum, d.ComplexNum,
	)
}

// ToString method of the Data structure converts all variable values into one string.
func (d *Data) ToString() string {
	return fmt.Sprintf("%d %o %X %g %s %t %v",
		d.NumDecimal, d.NumOctal, d.NumHexadecimal, d.Pi, d.Name, d.IsActive, d.ComplexNum)
}

// Hash method of the Data structure performs hashing of the string values with added salt.
func (d *Data) Hash() string {
	runes := []rune(d.ToString())
	salt := []rune(saltExample)
	mid := len(runes) / 2

	runes = append(runes[:mid], append(salt, runes[mid:]...)...)
	hash := sha256.Sum256([]byte(string(runes)))

	return fmt.Sprintf("%x", hash)
}
