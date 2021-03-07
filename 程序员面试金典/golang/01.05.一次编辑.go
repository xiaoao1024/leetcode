package main

/*
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。

示例 1:

输入:
first = "pale"
second = "ple"
输出: True


示例 2:

输入:
first = "pales"
second = "pal"
输出: False

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/one-away-lcci
*/

func oneEditAway(first string, second string) bool {
	s1, s2 := first, second
	if len(first) < len(second) {
		s1, s2 = second, first
	}

	// 比较长度
	if len(s1)-len(s2) > 1 {
		return false
	}

	var (
		j         int
		oneOp     bool
		isSameLen = len(s1) == len(s2)
	)

	for i := 0; i < len(s2); i++ {
		if s2[i] != s1[j] {
			if !oneOp {
				if isSameLen {
					// 1次update操作
				} else {
					// 1次delete或者add操作
					i--
				}
				oneOp = true
			} else {
				return false
			}
		}
		j++
	}

	return true
}
