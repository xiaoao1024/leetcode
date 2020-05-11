package solution

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		len1           = len(nums1)
		len2           = len(nums2)
		mid            = (len1 + len2) / 2
		index1, index2 int
		n1, n2         int
	)

	for i := 0; i < mid+1; i++ {
		n1 = n2
		if index1 < len1 && index2 < len2 {
			if nums1[index1] < nums2[index2] {
				n2 = nums1[index1]
				index1++
			} else {
				n2 = nums2[index2]
				index2++
			}
		} else if index1 < len1 {
			n2 = nums1[index1]
			index1++
		} else {
			n2 = nums2[index2]
			index2++
		}
	}

	if (len1+len2)%2 == 0 {
		return float64(n1+n2) / 2.0
	}

	return float64(n2)

}
