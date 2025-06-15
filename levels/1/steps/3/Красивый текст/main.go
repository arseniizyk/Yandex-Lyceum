package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var wordsCount, total int
	var pattern, text string
	fmt.Scan(&wordsCount, &pattern)
	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')

	pattern = strings.ToLower(pattern)
	text = strings.ToLower(text)

	words := strings.Fields(text)
	for _, word := range words {
		if word == pattern {
			total++
		}
	}

	fmt.Println(total)
}
