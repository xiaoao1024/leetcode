package solution

import (
	"testing"
)

func TestGetHint(t *testing.T) {
	res := getHint("1236", "1345")
	expect := "1A1B"
	if res != expect {
		t.Errorf("expect:%v\tactual:%v\n", expect, res)
	}
}
