package leetcode

func twoSum(nums []int, target int) []int {
	addendToIndex := make(map[int]int)
	for i, v := range nums {
		j, ok := addendToIndex[v]
		addendToIndex[target-v] = i
		if ok {
			return []int{j,i}
		}
	}
	return []int{}
}
