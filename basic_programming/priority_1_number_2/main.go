package main

import (
	"fmt"
)

func main() {
	var score int
	fmt.Print("Enter the score: ")
	fmt.Scanln(&score)

	switch {
	case score >= 85 && score <= 100:
		fmt.Println("Grade: A")
	case score >= 70 && score <= 84:
		fmt.Println("Grade: B")
	case score >= 55 && score <= 69:
		fmt.Println("Grade: C")
	case score >= 40 && score <= 54:
		fmt.Println("Grade: D")
	case score >= 0 && score <= 39:
		fmt.Println("Grade: E")
	default:
		fmt.Println("Invalid score")
	}
}
