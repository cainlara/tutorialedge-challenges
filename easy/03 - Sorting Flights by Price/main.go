package main

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// Solution implementation
type FlightsSlice []Flight

func (fs FlightsSlice) Len() int {
	return len(fs)
}

func (fs FlightsSlice) Swap(i, j int) {
	fs[i], fs[j] = fs[j], fs[i]
}

func (fs FlightsSlice) Less(i, j int) bool {
	return fs[i].Price < fs[j].Price
}

//End of soulution

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
	sort.Sort(FlightsSlice(flights))
	// implement
	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d\n", flight.Origin, flight.Destination, flight.Price)
	}
}

func main() {
	flights := []Flight{
		{Origin: "BGA", Destination: "SFO", Price: 500},
		{Origin: "BGA", Destination: "MIA", Price: 300},
		{Origin: "BGA", Destination: "YUL", Price: 800},
	}

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}
