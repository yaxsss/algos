package hashmap

import "strings"

func wordPattern(pattern string, s string) bool {
	ss := strings.Fields(s)
	if len(ss) != len(pattern) {
		return false
	}
	m := make(map[byte]string, len(ss))
	m2 := make(map[string]byte, len(ss))
	for i := range pattern {
		if b, ok := m2[ss[i]]; !ok {
			m2[ss[i]] = pattern[i]
		} else if b != pattern[i] {
			return false
		}
		if word, ok := m[pattern[i]]; !ok {
			m[pattern[i]] = ss[i]
		} else if word != ss[i] {
			return false
		}
	}
	return true
}
