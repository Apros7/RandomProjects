package main

import (
	"fmt"
)

func generics() {
	var slice []int = []int{1, 2, 3, 4}
	fmt.Println(sumSlice[int](slice))

	var sliceFloat []float32 = []float32{1, 2, 3, 4}
	fmt.Println(sumSlice[float32](sliceFloat))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
