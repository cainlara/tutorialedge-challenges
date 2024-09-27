package main

import (
	"fmt"
	"math"
)

type MyInt int

func (n *MyInt) IsArmstrong() bool {
	firstDigit := *n / 100
	remainder := *n % 100
	secondDigit := remainder / 10
	thirdDigit := remainder % 10

	sum := math.Pow(float64(firstDigit), 3) + math.Pow(float64(secondDigit), 3) + math.Pow(float64(thirdDigit), 3)

	return MyInt(sum) == *n
}

func main() {
	fmt.Println("Armstrong Numbers")

	var num1 MyInt = 407
	fmt.Println(num1.IsArmstrong())
}
