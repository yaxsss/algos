package twopointers

func getLowerAlphabet(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c
	}
	if c >= 'A' && c <= 'Z' {
		return c - 'A' + 'a'
	}
	if c >= '0' && c <= '9' {
		return c
	}
	return 0
}

func isPalindrome(s string) bool {
	head := 0
	tail := len(s) - 1
	var c1, c2 byte
	for head < tail {
		for head <= tail {
			c1 = getLowerAlphabet(s[head])
			if c1 > 0 {
				break
			}
			head++
		}
		for head <= tail {
			c2 = getLowerAlphabet(s[tail])
			if c2 > 0 {
				break
			}
			tail--
		}
		// 重叠后，有可能c1和c2其中一个是非法字符
		if c1 > 0 && c2 > 0 && c1 != c2 {
			return false
		}
		head++
		tail--
	}
	return true
}
