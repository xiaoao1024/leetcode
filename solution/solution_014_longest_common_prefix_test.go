package solution

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs1 := []string{"flower", "flow", "flight"}
	if longestCommonPrefix(strs1) != "fl" {
		t.Errorf("res != fl")
	}

	strs2 := []string{"dog", "racecar", "car"}
	if longestCommonPrefix(strs2) != "" {
		t.Errorf("res != ")
	}
}
