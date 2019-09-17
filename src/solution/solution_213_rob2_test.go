package solution

import (
	"fmt"
	"testing"
)

func Test_rob2(t *testing.T) {
	testCase := [][]int{
		[]int{2, 3, 2},
		[]int{1, 2, 3, 1},
	}

	for _, c := range testCase {

		fmt.Println(rob2(c))
	}
}
