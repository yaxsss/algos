package ranged

import "fmt"

func summaryRange(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	left := 0
	count := 1
	res := []string{}
	for i := 1; i < len(nums); i++ {
		if nums[left]+count != nums[i] {
			if count == 1 {
				res = append(res, fmt.Sprintf("%d", nums[left]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[left], nums[i-1]))
			}
			count = 1
			left = i
		} else {
			count++
		}
	}
	if count == 1 {
		res = append(res, fmt.Sprintf("%d", nums[left]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[left], nums[len(nums)-1]))
	}
	return res
}
