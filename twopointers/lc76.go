package twopointers

import "math"

func minWindow(s string, t string) string {
	needs := make(map[byte]int)
	for i := range t {
		needs[t[i]]++
	}
	windows := make(map[byte]int)
	left := 0
	valid := 0
	start, minLen := 0, math.MaxInt
	for right := range s {
		c := s[right]
		windows[c]++
		if windows[c] == needs[c] {
			valid++
		}
		// 当窗口满足条件时，尝试移动left，缩小合法的窗口范围，并记录到minLen中, 直到窗口又变为不合法，窗口右指针向后移
		for valid == len(needs) {
			if right-left < minLen {
				start = left
				minLen = right - left
			}
			c := s[left]
			left++
			if _, ok := needs[c]; ok {
				if windows[c] == needs[c] {
					valid--
				}
				windows[c]--
			}
		}
	}
	if minLen == math.MaxInt {
		return ""
	}
	return s[start : start+minLen+1]
}
