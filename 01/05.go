package main

import "fmt"

func filterLong(strs []string) []string {
	result := []string{}
	for _, s := range strs {
		if len(s) > 3 {
			result = append(result, s)
		}
	}
	return result
}
func main() {
	fmt.Println(filterLong([]string{"go", "java", "js"}))
}
