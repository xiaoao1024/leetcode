package solution

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	strNum := strconv.Itoa(x)

	var (
		left  int
		right = len(strNum) - 1
	)

	for left < right {
		if strNum[left] != strNum[right] {
			return false
		}

		left++ // NOTE: 自增、自减不能作为表达式使用，只能当做语句。
		right--
	}

	return true
}
