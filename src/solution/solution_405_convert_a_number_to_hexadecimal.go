package solution

import (
	"fmt"
)

/**
Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, two’s complement method is used.

Note:

All letters in hexadecimal (a-f) must be in lowercase.
The hexadecimal string must not contain extra leading 0s. If the number is zero, it is represented by a single zero character '0'; otherwise, the first character in the hexadecimal string will not be the zero character.
The given number is guaranteed to fit within the range of a 32-bit signed integer.
You must not use any method provided by the library which converts/formats the number to hex directly.
*/

const hexStr = "0123456789abcdef"

func toHex(num int) string {
	var res string
	for i := 0; i < 4; i++ {
		tmp := num >> (4 * uint(i)) & 0xf
		fmt.Printf("%b %b %s\n", tmp, num, string(hexStr[tmp]))
		res = string(hexStr)
	}
	return res
}
