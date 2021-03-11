package main

import (
	"flag"
	"fmt"
	"strings"
)

func printWords(words []string) {
	fmt.Println(strings.Join(words, " "))
}

func printSentences(sentences []string) {

	fmt.Println(strings.Join(sentences, ". "))
}

func main() {

	wordCount := flag.Int("w", 0, "Number of words")
	sentenceCount := flag.Int("s", 0, "Number of sentenceCount")
	flag.Parse()

	if *sentenceCount > 0 {
		sentences := getSentences(*sentenceCount)
		printSentences(sentences)
		return
	}

	if *wordCount > 0 {
		var words = getWords(*wordCount)
		printWords(words)
		return
	}

	flag.Usage()
}
