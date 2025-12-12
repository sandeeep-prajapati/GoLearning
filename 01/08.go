package main

import "fmt"

func main() {
	text := "hello World"
	freq := make(map[rune]int)

	for _, ch := range text {
		freq[ch]++
	}
	fmt.Println(freq)
}
