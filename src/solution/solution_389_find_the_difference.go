package solution

import (
	"fmt"
)

func FindTheDifference(s string, t string) byte {
	var sum1, sum2 int32
	for _, ch := range s {
		sum1 += ch
	}

	for _, ch := range t {
		sum2 += ch
	}

	return byte(sum2 - sum1)
}
