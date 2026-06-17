package main

import "fmt"

func TriangularNumber(n int) int {
	// implement me!
	return (n * (n + 1)) / 2
}

func main() {
	fmt.Println("Returning the 'nth' triangular number")

	number := TriangularNumber(3)
	fmt.Println(number) // '6'
}
