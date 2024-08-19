package main

// https://go.dev/tour/basics/7

import (
	"fmt"
	"math"
	"math/rand"
)

func arrays(length int) (x, y int) {
	x = length * 4 / 9
	y = length - x
	return
}

func prints() {
	fmt.Println("Your random number is:", rand.Intn(10))
	fmt.Println("Pi is:", math.Pi)
	fmt.Println(arrays(17))
}
