package solution

import (
	"testing"
)

func TestSumNums(t *testing.T) {
	if sumNums(9) != 45 {
		t.Errorf("sumNums(9) != 45")
	}
}
