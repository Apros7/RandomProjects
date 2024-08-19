package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	prints()
	var buff int = 0
	if buff == 1 {
		prints()
		test()
		generics()
		channels()
		concurrency()
		structs()
		strings()
		loops()
		maps()
		array()
		intDivProgram()
	}
}

func test() {
	fmt.Println("Your random number is:", rand.Intn(10))
	fmt.Println("Pi is:", math.Pi)
}
