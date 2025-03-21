package dp

func tribonacci(n int) int {
	t0, t1, t2 := 0, 1, 1
	if n == 0 {
		return t0
	}else if n == 1 {
		return t1	
	}else if n == 2 {
		return t2
	}

	for i := 3; i <= n; i++ {
		t0, t1, t2 = t1, t2, t0 + t1 + t2
	}
	return t2
}