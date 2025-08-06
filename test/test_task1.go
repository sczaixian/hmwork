package main

import (
	"fmt"
)

func two_sum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	valToIndex := make(map[int]int)
	for idx, num := range nums {
		complement := target - num
		//if (complement > 0 && num > 0 && target < 0) ||
		//	(complement < 0 && num < 0 && target > 0) { return nil}
		if tar, exists := valToIndex[complement]; exists {
			return []int{idx, tar}
		}
		valToIndex[num] = idx
	}
	return []int{}
}

func main() {
	fmt.Println("Hello World")
	//var res []int
	res := two_sum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}
