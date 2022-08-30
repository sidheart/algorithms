package leetcode

type tuple struct {
	Temperature, Index int
}

func dailyTemperatures(T []int) []int {
	stack := make([]tuple, len(T))
	answer := make ([]int, len(T))
	for i, temp := range T {
		for len(stack) > 0 && temp > stack[len(stack) - 1].Temperature {
			poppedTuple := stack[len(stack) - 1]
			answer[poppedTuple.Index] = i-poppedTuple.Index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, tuple{temp, i})
	}
	return answer
}
