package solution

func Intersection(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]bool, 0)
	map2 := make(map[int]bool, 0)
	for _, num := range nums1 {
		map1[num] = true
	}

	for _, num := range nums2 {
		map2[num] = true
	}

	res := make([]int, 0)
	for num, _ := range map1 {
		if _, ok := map2[num]; ok {
			res = append(res, num)
		}
	}
	return res
}
