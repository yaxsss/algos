package stringss

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	validPalindrome := func(left, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	maxLen := 1
	begin := 0

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i+1 > maxLen && validPalindrome(i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}

	maxLen := 1
	begin := 0

	expandAroundCenter := func(left, right int) int {
		i := left
		j := right
		for i >= 0 && j < len(s) {
			if s[i] == s[j] {
				i--
				j++
			} else {
				break
			}
		}
		// 跳出for循环时刚好s[i] != s[j]
		// 回文长度就是(j-1) - (i+1) + 1 = j-i-1
		return j - i - 1
	}

	for i := 0; i < len(s)-1; i++ {
		oddLen := expandAroundCenter(i, i)
		evenLen := expandAroundCenter(i, i+1)
		curMaxLen := max(oddLen, evenLen)
		if curMaxLen > maxLen {
			maxLen = curMaxLen
			// i是中心位置索引，maxLen是长度，求开始位置
			begin = i - (maxLen-1)/2
		}
	}
	return s[begin : begin+maxLen]
}
