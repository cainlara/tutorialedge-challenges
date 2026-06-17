package main

import (
	"errors"
	"fmt"
	"math"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

func GetMinMax(flights []Flight) (int, int, error) {
	if len(flights) == 0 {
		return -1, -1, errors.New("the slice is empty")
	}

	cheap := math.MaxInt
	expensive := math.MinInt

	for _, f := range flights {
		if f.Price < cheap {
			cheap = f.Price
		}

		if f.Price > expensive {
			expensive = f.Price
		}
	}

	return cheap, expensive, nil
}

func main() {
	flights := []Flight{
		{"A", "B", 400},
		{"B", "C", 300},
		{"C", "D", 200},
		{"D", "A", 100},
	}
	fmt.Println("Getting the Minimum and Maximum Flight Prices")
	cheapest, expensive, err := GetMinMax(flights)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Cheapest flight costs %d, the most expensive costs %d", cheapest, expensive)
}
