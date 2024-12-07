// This program demonstrate the entire pipeline process.
package main

import (
	"fmt"

	cv "github.com/kossadda/go-core-task/9/pkg/convey"
)

func main() {
	data := []uint8{1, 2, 3, 4, 5}

	cube := cv.Output(cv.SecondStep(cv.FirstStep(data)))

	for i := 0; i < len(data) && i < len(cube); i++ {
		fmt.Printf("(%d:%.0f) ", data[i], cube[i])
	}
	fmt.Println("")
}
