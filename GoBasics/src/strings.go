package main

import (
	"fmt"
)

func strings() {
	var simpleString string = "This is a string"
	fmt.Println(simpleString[0]) // Gives 84: Ascii value of "T", although this can fail if not an ascii character
	// If we use a for loop, it will give the correct value:

	for i, v := range "résumé" {
		fmt.Println(i, v)
	}
	// or fix it using runes:
	var runeString = []rune("résumé")
	fmt.Println(runeString)
}
