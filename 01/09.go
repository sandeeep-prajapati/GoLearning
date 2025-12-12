package main

import "fmt"

func swapMap(m map[int]string) map[string]int {
	result := make(map[string]int)
	for k, v := range m {
		result[v] = k
	}
	return result
}

func main() {
	m := map[int]string{1: "apple", 2: "banana"}
	fmt.Println(swapMap(m))
}
