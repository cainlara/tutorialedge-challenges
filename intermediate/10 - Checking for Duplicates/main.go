package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	mappedDevs := make(map[string]bool, len(developers))
	filtered := make([]string, 0, len(developers))

	for _, dev := range developers {
		if _, found := mappedDevs[dev.Name]; !found {
			filtered = append(filtered, dev.Name)
		}

		mappedDevs[dev.Name] = true
	}

	return filtered
}

func main() {
	fmt.Println("Filter Unique Challenge")
	devs := []Developer{
		{Name: "Elliot"},
		{Name: "Alan"},
		{Name: "Jennifer"},
		{Name: "Graham"},
		{Name: "Paul"},
		{Name: "Alan"},
	}

	filtered := FilterUnique(devs)
	fmt.Println(filtered)
}
