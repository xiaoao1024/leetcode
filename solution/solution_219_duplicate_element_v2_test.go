package solution

import (
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	arr_nums := [][]int{
		[]int{1, 2, 3, 1},
		[]int{1, 0, 1, 1},
		[]int{1, 2, 3, 1, 2, 3},
	}

	arr_k := []int{3, 1, 2}

	arr_resut := []bool{true, true, false}

	for i := 0; i < len(arr_nums); i++ {
		res := containsNearbyDuplicate(arr_nums[i], arr_k[i])
		if res != arr_resut[i] {
			t.Errorf("input nums:%v k:%v, expect:%v, but:%v", arr_nums[i], arr_k[i], arr_resut[i], res)
		}
	}
}
