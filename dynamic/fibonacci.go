package dynamic

// Fib computes the nth fibonacci number
func Fib(n int) int {
	fibber := fib()
	return fibber(n)
}

// fib is a closure that computes the nth fiboniacci number
func fib() func(int) int {
	memo := []int{0, 1}
	return func(n int) int {
		for len(memo) < n+1 {
			length := len(memo)
			memo = append(memo, memo[length-1]+memo[length-2])
		}
		return memo[n]
	}
}
