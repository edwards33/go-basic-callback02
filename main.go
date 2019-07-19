package main

import (
	"fmt"
)

func filter(nums []int, cb func(int) bool) []int{
	xs := []int{}

	for _, n := range nums {
		if cb(n) {
			xs = append(xs, n)
		}
	}

	return xs
}

func main() {
	xs := filter([]int{3,5,1,1,7}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)
}