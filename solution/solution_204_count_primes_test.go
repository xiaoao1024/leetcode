package solution

import (
	"testing"
)

func Test_countPrimes(t *testing.T) {
	testCases := []int{10}
	expectResults := []int{4}

	for i := 0; i < len(testCases); i++ {
		if countPrimes(testCases[i]) != expectResults[i] {
			t.Errorf("input:%d, expect:%d, but:%d\n", testCases[i], expectResults[i], countPrimes(testCases[i]))
		}
	}

}
