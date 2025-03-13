package twopointers

func findSubString(s string, words []string) []int {
	ws := map[string]int{}
	for _, word := range words {
		ws[word]++
	}
	num := len(words)
	span := len(words[0])

	arr := []int{}
	for k := 0; k < span && k+num*span <= len(s); k++ {
		valid := 0
		windows := map[string]int{}
		left := k
		for i := k; i+span <= len(s); i += span {
			word := s[i : i+span]
			windows[word]++
			// 计算窗口新增的数据是否合法：
			// 1. 不合法：新加的word不存在或超限, 移动左指针, 如果当前word根据不在words中，则跳过left跳过该word，否则left就往前移，直到word个数不超限
			// 2. 合法：将word valid数量新增, 当word valid的数量等于总数量，则认为这个窗口达到要求。加入记录, 然后移动左指针，进行新的窗口尝试
			if _, ok := ws[word]; ok && windows[word] <= ws[word] {
				valid++
				if valid == num {
					arr = append(arr, left)
					valid--
					windows[s[left:left+span]]--
					left += span
				}
			} else if !ok {
				valid = 0
				left = i + span
				windows = map[string]int{}
			} else {
				for windows[word] > ws[word] {
					windows[s[left:left+span]]--
					left += span
					valid--
				}
				valid++
			}
		}
	}
	return arr
}
