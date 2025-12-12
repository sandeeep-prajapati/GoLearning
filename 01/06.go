package main

import "fmt"

func doubleSlice(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}
func main() {
	arr := []int{1, 2, 3}
	doubleSlice(arr)
	fmt.Println(arr)
}
