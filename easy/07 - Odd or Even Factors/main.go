package main

import (
	"fmt"
)

func OddEvenFactors(num int) string {
	factors := findFactors(num)
	if factors%2 == 0 {
		return "even"
	}
	return "odd"
}

func findFactors(num int) int {
	counter := 0

	for div := 2; div <= num/2; div++ {
		if num%div == 0 {
			counter++
		}
	}

	return counter + 2 //Include "1" and num itself as factors
}

func main() {
	fmt.Println("Odd or Even Factors")

	numFactors := OddEvenFactors(23)
	fmt.Println(numFactors) // "even"

	numFactors = OddEvenFactors(36)
	fmt.Println(numFactors) // "odd"
}
