package solution

import (
	"fmt"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	testCase := []int{1, 2, 3, 4, 5, 6}
	for _, c := range testCase {
		fmt.Println(climbStairs(c))
	}
}
