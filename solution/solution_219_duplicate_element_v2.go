package solution

/*
给定一个整数数组，判断是否存在重复元素。

如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。
*/

func containsDuplicate(nums []int) bool {
	map1 := make(map[int]struct{}, 0)

	for _, num := range nums {
		if _, ok := map1[num]; ok {
			return true
		} else {
			map1[num] = struct{}{}
		}
	}

	return false
}
