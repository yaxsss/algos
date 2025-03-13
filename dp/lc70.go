package dp

func climbStars(n int) int {
	steps := make([]int, n+1)
	steps[0] = 1
	steps[1] = 1
	for i := 2; i < n+1; i++ {
		steps[i] = steps[i-1] + steps[i-2]
	}
	return steps[n]
}

func climbStars2(n int) int {
	prev, next := 1, 1
	// 第n阶和第n+1阶之和就是第n+2阶的步数
	// 每次计算一阶，保存在next中，prev中保存上一阶的值
	for i := 2; i <= n; i++ {
		next, prev = prev+next, next
	}
	return next
}
