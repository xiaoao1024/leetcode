package solution

import (
	"fmt"
	"testing"
)

func Test_rob(t *testing.T) {
	testCase := [][]int{
		[]int{1, 2, 3, 1},
		[]int{2, 7, 9, 3, 1},
	}

	for _, c := range testCase {

		fmt.Println(rob(c))
	}
}
