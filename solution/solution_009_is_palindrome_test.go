package solution

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	testCases := []int{121, -121, 10}
	expectResults := []bool{true, false, false}

	for i, in := range testCases {
		res := isPalindrome(in)
		if res != expectResults[i] {
			t.Errorf("input:%d, expect:%v, but:%v", in, expectResults[i], res)
		}
	}
}
