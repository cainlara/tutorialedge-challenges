package main

import "fmt"

func CheckPermutations(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	mappedStr1 := make(map[rune]int, len(str1))
	for _, c := range str1 {
		if count, found := mappedStr1[c]; found {
			mappedStr1[c] = count + 1
		} else {
			mappedStr1[c] = 1
		}
	}

	mappedStr2 := make(map[rune]int, len(str2))
	for _, c := range str2 {
		if count, found := mappedStr2[c]; found {
			mappedStr2[c] = count + 1
		} else {
			mappedStr2[c] = 1
		}
	}

	for key, value := range mappedStr1 {
		if mappedStr2[key] != value {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Check Permutations Challenge")

	str1 := "adcme"
	str2 := "medac"

	if CheckPermutations(str1, str2) {
		fmt.Printf("%s is a permutation of %s\n", str2, str1)
	} else {
		fmt.Printf("%s is not a permutation of %s\n", str2, str1)
	}
}
