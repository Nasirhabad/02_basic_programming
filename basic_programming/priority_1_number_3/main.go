package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i * i)
		} else {
			fmt.Println(i)
		}
	}

	// If the number is divisible by 7 then print the power of 3 value of the number.

	for i := 1; i <= 100; i++ {
		if i%7 == 0 {
			fmt.Printf("%d: %.0f\n", i, math.Pow(float64(i), 3))
		} else {
			fmt.Println(i)
		}
	}

	// If the number is divisible by 4 and 7 then print “buzz”.

	for i := 1; i <= 100; i++ {
		if i%4 == 0 && i%7 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
