package backtrack

var nums = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinationsBacktrack(path []byte, digits string, index int, result *[]string) {
	chars := nums[digits[index]]
	for _, c := range chars {
		path[index] = c
		if index == len(digits)-1 {
			*result = append(*result, string(path))
		} else {
			letterCombinationsBacktrack(path, digits, index+1, result)
		}
	}
}

func letterCombinations(digits string) []string {
	var result []string
	path := make([]byte, len(digits))
	letterCombinationsBacktrack(path, digits, 0, &result)
	return result
}
