package hashmap

func canConstruct(ransomNote string, magazine string) bool {
	chars := make(map[byte]int)

	for i := range magazine {
		c := magazine[i]
		chars[c]++
	}

	for i := range ransomNote {
		c := ransomNote[i]
		if i, ok := chars[c]; !ok || i == 0 {
			return false
		}
		chars[c]--
	}
	return true

}
