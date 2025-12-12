package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 1; i <= 5; i++ {
		fmt.Println(<-results)
	}
}
