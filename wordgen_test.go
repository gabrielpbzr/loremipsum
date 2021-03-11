package main

import "testing"

// Test word generation
func TestShouldGenerateFiveWords(t *testing.T) {
	count := 5
	words := getWords(count)
	if len(words) != count {
		t.Errorf("Should have generated %d words", count)
	}
}

// Test sentence generation
func TestShouldGenerateFiveSentences(t *testing.T) {
	count := 5
	sentences := getSentences(count)
	if len(sentences) != count {
		t.Errorf("Should have generated %d sentences", count)
	}
}
