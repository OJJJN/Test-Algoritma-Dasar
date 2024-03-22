package main

import "fmt"

func main() {
	// Membuat array
	arr := []interface{}{12, 9, 30, "A", "M", 99, 82, "J", "N", "B"}
	sortArray(arr)
	fmt.Println("Sorted Array:", arr)

	// Print pola
	fmt.Println("Pattern count:", patternCount("palindrom", "ind"))
	fmt.Println("Pattern count:", patternCount("ababab", "aba"))
	fmt.Println("Pattern count:", patternCount("abakadabra", "ab"))
	fmt.Println("Pattern count:", patternCount("aaaaaa", "aa"))
	fmt.Println("Pattern count:", patternCount("hello", ""))
	fmt.Println("Pattern count:", patternCount("hell", "hello"))

	// Print huruf
	fmt.Println("Count characters:", countCharacters("Hello World"))
	fmt.Println("Count characters:", countCharacters("Bismillah"))
	fmt.Println("Count characters:", countCharacters("MasyaAllah"))
}
