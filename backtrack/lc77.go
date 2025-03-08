package backtrack

func combineBacktrack(path []int, start, pos int, n, k int, result *[][]int) {
	for i := start; i <= n-k+pos; i++ {
		path[pos-1] = i
		if pos == k {
			r := make([]int, len(path))
			copy(r, path)
			*result = append(*result, r)
		} else {
			combineBacktrack(path, i+1, pos+1, n, k, result)
		}
	}
}
func combine(n int, k int) [][]int {
	path := make([]int, k)
	result := [][]int{}
	combineBacktrack(path, 1, 1, n, k, &result)
	return result
}
