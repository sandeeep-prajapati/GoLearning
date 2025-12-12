package main

import (
	"fmt"
	"time"
)

func printMsg() {
	fmt.Println("hello from goroutine!")
}

func main() {
	go printMsg()
	time.Sleep(time.Second)
}
