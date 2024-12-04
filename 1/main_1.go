package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := Data{
		numDecimal:     42,
		numOctal:       052,
		numHexadecimal: 0x2A,
		pi:             3.14,
		name:           "Golang",
		isActive:       true,
		complexNum:     complex(1.0, 2.0),
	}

	fmt.Println(&data)
	fmt.Println("Sha256 -", data.Hash())
}

type Data struct {
	numDecimal     int
	numOctal       int
	numHexadecimal int
	pi             float64
	name           string
	isActive       bool
	complexNum     complex64
}

func (d *Data) String() string {
	return fmt.Sprintf("numDecimal %d - %T\nnumOctal %o - %T\nnumHexadecimal %X - %T\npi %g - %T\nname %s - %T\nisActive %t - %T\ncomplexNum %v - %T",
		d.numDecimal, d.numDecimal, d.numOctal, d.numOctal, d.numHexadecimal, d.numHexadecimal, d.pi, d.pi, d.name, d.name, d.isActive, d.isActive, d.complexNum, d.complexNum)
}

func (d *Data) ToString() string {
	return fmt.Sprintf("%d %o %X %g %s %t %v", d.numDecimal, d.numOctal, d.numHexadecimal, d.pi, d.name, d.isActive, d.complexNum)
}

func (d *Data) Hash() string {
	runes := []rune(d.ToString())
	salt := []rune("go-2024")
	mid := len(runes) / 2

	runes = append(runes[:mid], append(salt, runes[mid:]...)...)
	hash := sha256.Sum256([]byte(string(runes)))

	return fmt.Sprintf("%x", hash)
}
