package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Flight struct which contains
// the origin, destination and price of a flight
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// IsSubset checks to see if the first set of
// flights is a subset of the second set of flights.
func IsSubset(first, second []Flight) bool {
	mappedSecond := make(map[string]int, len(second))

	for _, f := range second {
		hash := Hash(f)
		if amount, found := mappedSecond[hash]; found {
			mappedSecond[hash] = amount + 1
		} else {
			mappedSecond[hash] = 1
		}
	}

	for _, f := range first {
		hash := Hash(f)

		if _, found := mappedSecond[hash]; !found {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Sets and Subsets Challenge")
	firstFlights := []Flight{
		{Origin: "GLA", Destination: "CDG", Price: 1000},
		{Origin: "GLA", Destination: "JFK", Price: 5000},
		{Origin: "GLA", Destination: "SNG", Price: 3000},
	}

	secondFlights := []Flight{
		{Origin: "GLA", Destination: "CDG", Price: 1000},
		{Origin: "GLA", Destination: "JFK", Price: 5000},
		{Origin: "GLA", Destination: "SNG", Price: 3000},
		{Origin: "GLA", Destination: "AMS", Price: 500},
	}

	subset := IsSubset(firstFlights, secondFlights)
	fmt.Println(subset)
}

func Hash(f Flight) string {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(f)
	return b.String()
}
