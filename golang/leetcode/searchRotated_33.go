package leetcode

// Returns the rotated index of an array that was rotated right by rot places
func rotate(i, rot, n int) int {
	return (i + rot) % n
}

// Returns the index of the target value in a sorted array that may be rotated by n places
// or returns -1 if target is not found in the array
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	rot := hi
	lo, hi = 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		rotatedMid := rotate(mid, rot, len(nums))
		if nums[rotatedMid] == target {
			return rotatedMid
		} else if nums[rotatedMid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
