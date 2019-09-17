package solution

import (
	"fmt"
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
	testCase := [][]string{
		[]string{"abc", "def"},
		[]string{"aaa", "bbb"},
		[]string{"abbcc", "bbbac"},
		[]string{"aabbcc", "bbeecc"},
	}

	for _, c := range testCase {
		fmt.Println(isIsomorphic(c[0], c[1]))
	}
}
