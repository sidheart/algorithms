package main

import (
	"math"
)

func checkPossibility(nums []int) bool {
	invariants := 0
	min := math.MinInt64
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i+1] {
			invariants += 1
			if min > nums[i+1] {
				nums[i+1] = nums[i]
			}
		} else {
			min = nums[i]
		}
	}
	return invariants <= 1
}
