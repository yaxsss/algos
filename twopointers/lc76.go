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
