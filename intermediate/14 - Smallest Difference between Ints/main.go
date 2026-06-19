package main

import (
	"fmt"
	"sort"
)

func CalcSmallestDifference(arr1, arr2 []int) int {
	sort.Ints(arr1)
	sort.Ints(arr2)

	return arr2[0] - arr1[len(arr1)-1]
}

func main() {
	fmt.Println("Smallest Difference Challenge")

	arr1 := []int{9, 8, 7, 6}
	arr2 := []int{7, 3, 2, 5}

	smallestDiff := CalcSmallestDifference(arr1, arr2)
	fmt.Println(smallestDiff)
}
