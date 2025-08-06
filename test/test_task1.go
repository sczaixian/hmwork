package main

import (
	"fmt"
	"sort"
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

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(x, y int) bool {
		return intervals[x][0] < intervals[y][0]
	})

	res := [][]int{}
	tmp := intervals[0]

	for _, item := range intervals[1:] {
		if tmp[1] > item[0] && item[1] > tmp[1] {
			tmp[1] = item[1]
		} else {
			res = append(res, tmp)
			tmp = item
		}
	}
	res = append(res, tmp)
	return res
}

func removeDuplicates(nums []int) int {
	if len(nums < 2) {
		return len(nums)
	}
	l := 0
	for _, num := range nums[1:] {
		if nums[l] != num {
			l += 1
			nums[l] = num
		}
	}
	return l + 1
}

func plusOne(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i] += 1
			return nums
		}
		nums[i] = 0
	}
	return append([]int{1}, nums...) //
}

func main() {
	fmt.Println("Hello World")
	res := two_sum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}
