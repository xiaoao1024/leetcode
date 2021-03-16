package main

import (
	"fmt"
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

func compressString(S string) string {
	if len(S) < 2 {
		return S
	}

	var (
		arr = make([]rune, 0)
		t   rune
		cnt int
	)

	for i, ch := range S {
		if i == 0 {
			t = ch
			cnt = 1
		} else {
			if ch == t {
				cnt++
			} else {
				arr = append(arr, t)
				arr = append(arr, []rune(fmt.Sprint(cnt))...)

				if len(arr) >= len(S) {
					return S
				}

				t = ch
				cnt = 1
			}
		}
	}

	arr = append(arr, t)
	arr = append(arr, []rune(fmt.Sprint(cnt))...)
	if len(arr) >= len(S) {
		return S
	}

	return string(arr)
}
