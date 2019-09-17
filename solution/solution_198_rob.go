package solution

func rob(nums []int) int {
	var pre2, pre1, now int
	for _, num := range nums {
		now = maxInt(pre2+num, pre1)
		pre2 = pre1
		pre1 = now
	}

	return now
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
