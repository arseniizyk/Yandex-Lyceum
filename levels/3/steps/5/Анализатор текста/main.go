package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func AnalyzeText(text string) {
	text = strings.ToLower(text)
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	wordMap := make(map[string]int)
	for _, w := range words {
		wordMap[w]++
	}

	mostPopularWord, mostPopularCount := mostPopular(text)

	fmt.Printf("Количество слов: %d\n", len(words))
	fmt.Printf("Количество уникальных слов: %d\n", len(wordMap))
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", mostPopularWord, mostPopularCount)

	topWords := getTopWords(wordMap, 5)

	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for _, w := range topWords {
		fmt.Printf("\"%s\": %d раз\n", w, wordMap[w])
	}

}

func mostPopular(text string) (string, int) {
	var mostPopular string
	var count int
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	for _, w := range words {
		if i := strings.Count(text, w); i > count {
			count = i
			mostPopular = w
		}
	}

	return mostPopular, count

}

func getTopWords(wordMap map[string]int, n int) []string {
	keys := make([]string, 0)
	for key := range wordMap {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return wordMap[keys[i]] > wordMap[keys[j]]
	})

	return keys[:n]
}
