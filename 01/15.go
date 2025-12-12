package main

import "fmt"

type Sortable interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

func BubbleSort(s Sortable) {
	n := s.Len()
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if s.Less(j+1, j) {
				s.Swap(j, j+1)
			}
		}
	}
}

type Numbers []int

func (n Numbers) Len() int           { return len(n) }
func (n Numbers) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Numbers) Less(i, j int) bool { return n[i] < n[j] }

func main() {
	nums := Numbers{5, 1, 4, 2, 3}
	BubbleSort(nums)
	fmt.Println(nums)
}
