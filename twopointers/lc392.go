package twopointers

func isSubsequence(s string, t string) bool {
	j := 0
	for i := range s {
		if s[i] == t[j] {
			j++
		}
		if j >= len(t) {
			return true
		}
	}
	return false
}
