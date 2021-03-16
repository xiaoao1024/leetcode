package main

import (
	"testing"
)

/*
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

示例1:

 输入："aabcccccaaa"
 输出："a2b1c5a3"
示例2:

 输入："abbccd"
 输出："abbccd"
 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
提示：

字符串长度在[0, 50000]范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/compress-string-lcci
*/

func Test_compressString(t *testing.T) {
	testCases := [][]string{
		{"aabcccccaaa", "a2b1c5a3"},
		{"abbccd", "abbccd"},
		{"rrrrrLLLLLPPPPPPRRRRRgggNNNNNVVVVVVVVVVDDDDDDDDDDIIIIIIIIIIlllllllAAAAqqqqqqqbbbNNNNffffff", "r5L5P6R5g3N5V10D10I10l7A4q7b3N4f6"},
	}

	var ret string
	for _, item := range testCases {
		ret = compressString(item[0])
		if ret != item[1] {
			t.Errorf("input:%v, output:%v", item, ret)
		}
	}
}
