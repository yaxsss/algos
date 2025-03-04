package array

// 先根据k位置，依次对左边倒序，然后对右边倒序，然后对整个结构倒序
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	swap := func(a, b *int) {
		tmp := *a
		*a = *b
		*b = tmp
	}
	len := len(nums)
	midLeft := (len - k) / 2
	right := len - k - 1
	for i := 0; i < midLeft; i++ {
		swap(&nums[i], &nums[right-i])
	}

	midRight := (len - k) + k/2
	right = len - 1
	for i := len - k; i < midRight; i++ {
		swap(&nums[i], &nums[right])
		right--
	}

	for i := 0; i < len/2; i++ {
		swap(&nums[i], &nums[len-i-1])
	}
}
