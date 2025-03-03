package array

func merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	for i := 0; i < n; i++ {
		for nums1[j] <= nums2[i] && j < m+i {
			j++
		}
		for k := m + i; k > j; k-- {
			nums1[k] = nums1[k-1]
		}
		nums1[j] = nums2[i]
	}
}

// TODO: 尝试设计复杂度O(m+n)
