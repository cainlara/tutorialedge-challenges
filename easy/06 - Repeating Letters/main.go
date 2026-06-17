package main

import (
	"fmt"
)

func DoubleChars(original string) string {
	runes := make([]rune, len(original)*2)
	runeIndex := 0

	for i := range len(original) {
		runes[runeIndex] = rune(original[i])
		runes[runeIndex+1] = rune(original[i])

		runeIndex = runeIndex + 2
	}

	return string(runes)
}

func main() {
	original := "gophers"
	doubled := DoubleChars(original)
	fmt.Println(doubled) // ggoopphheerrss
}
