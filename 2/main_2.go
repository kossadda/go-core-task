package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Slice []int

const (
	sliceSize        = 10
	maxValue         = 1000
	minValue         = -1000
	conditionCapSize = 1 << 10
	conditionCapExt  = 1.25
)

func main() {
	originalSlice := generateSlice(sliceSize)
	evenSlice := sliceExample(originalSlice)
	appendedSlice := addElements(originalSlice, 2345)

	fmt.Println("originalSlice:", originalSlice)
	fmt.Println("sliceExample:", evenSlice)
	fmt.Println("originalSlice:", appendedSlice, "added 2345")
}

func generateSlice(n int) Slice {
	src := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(src)

	slice := make(Slice, n)

	for i := 0; i < n; i++ {
		slice[i] = randGen.Intn(maxValue-minValue) + minValue
	}

	return slice
}

func sliceExample(slice Slice) Slice {
	var evenSlice Slice

	for _, num := range slice {
		if num%2 == 0 {
			evenSlice = append(evenSlice, num)
		}
	}

	return evenSlice
}

func addElements(slice Slice, val int) (s Slice) {
	l := len(slice)
	newCap := func() int {
		if len(slice) < conditionCapSize {
			return cap(slice) * 2
		} else {
			return int(float64(cap(slice)) * conditionCapExt)
		}
	}

	if l+1 < cap(slice) {
		s = slice[:l+1]
	} else {
		s = make(Slice, l+1, newCap())
	}

	copy(s, slice)
	s[l] = val

	return s
}
