package main

import (
	"errors"
	"fmt"
)

func intDivProgram() {
	const welcomeStatement string = "Welcome to the division program"
	var div1 int
	var div2 int
	fmt.Println(welcomeStatement)
	fmt.Println("Please input the nominator")
	fmt.Scanln(&div1)
	fmt.Println("Please input the denominator")
	fmt.Scanln(&div2)
	var result, remainder, err = intDiv(div1, div2)
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result is a complete division with value %v", result)
	default:
		fmt.Printf("The result is a division with value %v and remainder %v", result, remainder)
	}
	fmt.Printf("\n")
}

func intDiv(nom int, denom int) (int, int, error) {
	var err error
	if denom == 0 {
		err = errors.New("Cannot divide by Zero")
		return 0, 0, err
	}
	var division int = nom / denom
	var remainder int = nom % denom
	return division, remainder, err
}
