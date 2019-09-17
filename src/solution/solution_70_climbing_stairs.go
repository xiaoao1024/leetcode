package solution

func climbStairs(n int) int {
	if n < 1 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	var (
		pre2 = 1
		pre1 = 2
		now  int
	)

	for i := 2; i < n; i++ {
		now = pre2 + pre1
		pre2, pre1 = pre1, now
	}

	return now
}
