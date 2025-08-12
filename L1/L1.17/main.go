package main

import "fmt"

func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l := 0
	r := len(nums)
	for l+1 < r {
		m := (l + r) / 2
		if nums[m] > target {
			r = m
		} else {
			l = m
		}
	}
	if nums[l] == target {
		return l
	} else {
		return -1
	}
}

func main() {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
	}
	for _, test := range tests {
		fmt.Printf("nums: %v, target: %d, get: %d\n", test.nums, test.target, BinarySearch(test.nums, test.target))
	}
}
