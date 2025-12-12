package main

import "fmt"

func main() {
	students := map[string]int{
		"Sandeep": 100,
		"Aman":    1,
		"john":    75,
	}
	delete(students, "John")
	topper := ""
	topMarks := -1
	for name, marks := range students {
		if marks > topMarks {
			topMarks = marks
			topper = name
		}

	}
	fmt.Println("Topper is : ", topper)
}
