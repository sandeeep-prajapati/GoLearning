package main

import "fmt"

func increment(x *int) {
	*x += 10
}

func main() {
	a := 5
	increment(&a)
	fmt.Println(a)
}
