package solution

import (
	"testing"
)

func Test_isHappy(t *testing.T) {
	testCases := []int{7, 19, 24, 30}
	expectResults := []bool{true, true, false, false}

	for i, c := range testCases {
		if isHappy(c) != expectResults[i] {
			t.Errorf("input:%v, expect:%v, but:%v\n", c, expectResults[i], isHappy(c))
		}
	}
}
