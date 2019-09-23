package solution

import "testing"

func Test_containsDuplicate(t *testing.T) {
	testCases := [][]int{
		[]int{1, 2, 3, 1},
		[]int{1, 2, 3, 4},
		[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}

	expectResults := []bool{true, false, true}

	for i, nums := range testCases {
		res := containsDuplicate(nums)
		if res != expectResults[i] {
			t.Errorf("input:%v, expect:%v, but:%v", nums, expectResults[i], res)
		}
	}
}
