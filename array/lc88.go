package array

func merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	// 对nums2中的每个元素做操作：检查在nums1中第一个大于nums2的元素位置，将此位置开始的元素依次后移，腾出此位置给nums2的元素，
	// 最坏复杂度：O(mxn)
	for i := 0; i < n; i++ {
		// m+i代表这轮nums1中的最后一个元素
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
