package solution

/*
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	var (
		map1    = make(map[int]int, 0)
		minDist = len(nums)
	)

	for i, num := range nums {
		if index, ok := map1[num]; ok {
			dist := i - index
			if dist < minDist {
				minDist = dist
			}
		}
		map1[num] = i
	}

	if minDist == len(nums) {
		return false
	}

	return minDist <= k
}
