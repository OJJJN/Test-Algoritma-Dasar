package main

import (
	"strings"
	"unicode"
)

// melakukan hitung pola
func patternCount(text, pattern string) int {
	count := 0
	textLen := len(text)
	patternLen := len(pattern)

	for i := 0; i <= textLen-patternLen; i++ {
		if text[i:i+patternLen] == pattern {
			count++
		}
	}

	return count
}

// melakukan hitung huruf
func countCharacters(input string) map[string]int {
	result := make(map[string]int)

	for _, char := range input {
		if unicode.IsLetter(char) {
			charString := string(char)
			if unicode.IsLower(char) {
				result[charString]++
			} else {
				result[charString]++
				result[strings.ToLower(charString)]++
			}
		}
	}

	return result
}
