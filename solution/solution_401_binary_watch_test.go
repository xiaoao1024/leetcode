package solution

import (
	"fmt"
	"testing"
)

func Test_ReadBinaryWatch(t *testing.T) {
	for i := 0; i <= 10; i++ {
		res := readBinaryWatch(i)
		if len(res) > 0 {
			fmt.Printf("num=%d res=%v\n", i, res)
			fmt.Println()
		}
	}
}
