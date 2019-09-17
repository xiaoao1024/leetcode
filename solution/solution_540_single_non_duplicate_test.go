package solution

import (
	"fmt"
	"testing"
)

func Test_singleNonDuplicate(t *testing.T) {
	testCase := [][]int{
		[]int{1, 2, 2, 3, 3},
		[]int{1, 1, 2, 3, 3},
		[]int{1, 1, 2, 2, 3, 3, 4, 5, 5},
	}

	for _, c := range testCase {
		fmt.Println(singleNonDuplicate(c))
	}
}
