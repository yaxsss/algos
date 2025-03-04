package twopointers

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}
	last := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		// 相同的两个元素，得到的结果是一样的，直接忽略掉
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[j] + nums[k]
			if sum > -nums[i] {
				k--
			} else if sum < -nums[i] {
				j++
			} else {
				r := []int{nums[i], nums[j], nums[k]}
				// j，k前后元素相同,过滤掉【
				if len(last) == 0 || last[0] != r[0] || last[1] != r[1] {
					res = append(res, r)
					last = r
				}
				k--
				j++
			}
		}
	}
	return res
}
