package main

import "fmt"

func multiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func main() {
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(10))
	fmt.Println(triple(10))
}
