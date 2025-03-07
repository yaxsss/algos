package hashmap

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	srcs := make(map[byte]byte)
	dsts := make(map[byte]byte)
	for i := range s {
		if v, ok := srcs[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else {
			srcs[s[i]] = t[i]
		}

		if v, ok := dsts[t[i]]; ok {
			if v != s[i] {
				return false
			}
		} else {
			dsts[t[i]] = s[i]
		}
	}

	return true
}
