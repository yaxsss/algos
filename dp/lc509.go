package dp

func fib(n int) int {
	prev, next := 0, 1
	if n == 0 {
		return prev
	}
	for i := 2; i <= n; i++ {
		next, prev = prev+next, next
	}
	return next
}
