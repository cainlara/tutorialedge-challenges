package main

import (
	"fmt"
	"math"
)

func DiffSquares(n, m int) int {
	return int(math.Pow(float64(n), 2)) - int(math.Pow(float64(m), 2))
}

func main() {
	fmt.Println("Calculate The Difference of Squares")
	fmt.Println(DiffSquares(5, 4))
}
