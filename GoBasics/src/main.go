package main

import "fmt"

func main() {
	var intNum uint16 = 32767
	var floatNum float32 = 12345678.9
	fmt.Println(intNum)
	fmt.Println(floatNum)

	var intNum2 uint16 = 10000
	fmt.Println(intNum + intNum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	const myRune rune = 'a'
	fmt.Println(myRune)

	int1, int2 := 2, 2
	printMe(int1 + int2)

}

func printMe(param int) {
	fmt.Println(param)
}
