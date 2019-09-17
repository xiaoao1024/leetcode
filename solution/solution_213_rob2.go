package solution

func rob2(nums []int) int {
	n := len(nums)
	if nums == nil || n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	return maxInt(robWithRange(nums, 0, n-2), robWithRange(nums, 1, n-1))
}

func robWithRange(nums []int, first, last int) int {
	var pre2, pre1, now int
	for i := first; i <= last; i++ {
		now = maxInt(pre2+nums[i], pre1)
		pre2, pre1 = pre1, now
	}

	return now
}
