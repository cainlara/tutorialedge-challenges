package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type Word struct {
	Copy        string
	Occurrences int
}

func CountWords(text string) map[string]int {
	count := make(map[string]int)
	chars := []rune{}

	for _, char := range text {
		if !unicode.IsPunct(char) || char == '\'' {
			chars = append(chars, char)
		}
	}

	cleaned := string(chars)

	split := strings.Split(cleaned, " ")

	for _, word := range split {
		count[word]++
	}

	return count
}

func Top5Words(wordmap map[string]int) []Word {
	var allOcurrences []Word

	for copy, occurrences := range wordmap {
		allOcurrences = append(allOcurrences, Word{
			copy, occurrences,
		})
	}

	sort.Slice(allOcurrences, func(i, j int) bool {
		return allOcurrences[i].Occurrences > allOcurrences[j].Occurrences
	})

	topFive := make([]Word, 5)

	for index := 0; index < 5 && index < len(allOcurrences); index++ {
		topFive[index] = allOcurrences[index]
	}

	return topFive
}

func main() {
	fmt.Println("Word Frequency Test")

	text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`
	// text := `A A A A A A B B B B B C C C C D D D E E F`

	results := CountWords(text)
	MostCommon := Top5Words(results)

	fmt.Println(MostCommon)
}
