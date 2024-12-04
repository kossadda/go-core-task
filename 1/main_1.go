package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := Data{
		int:       42,
		float64:   3.14,
		string:    "Golang",
		bool:      true,
		complex64: complex(1.0, 2.0),
	}

	fmt.Println(&data)
	fmt.Println(data.Hash())
}

type Data struct {
	int
	float64
	string
	bool
	complex64
}

func (d *Data) String() string {
	return fmt.Sprintf("%T\n%T\n%T\n%T\n%T", d.int, d.float64, d.string, d.bool, d.complex64)
}

func (d *Data) ToString() string {
	return fmt.Sprintf("%d %o %X %f %s %t %v", d.int, d.int, d.int, d.float64, d.string, d.bool, d.complex64)
}

func (d *Data) Hash() string {
	runes := []rune(d.ToString())
	salt := []rune("go-2024")
	mid := len(runes) / 2

	runes = append(runes[:mid], append(salt, runes[mid:]...)...)
	hash := sha256.Sum256([]byte(string(runes)))

	return fmt.Sprintf("%x", hash)
}
