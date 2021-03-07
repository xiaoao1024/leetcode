package main

import (
	"testing"
)

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

func Test_oneEditAway(t *testing.T) {
	testCases := [][]interface{}{
		{"aple", "ple", true},
		{"aples", "pal", false},
		{"ab", "bc", false},
		{"teacher", "bleacher", false},
	}

	var ret bool
	for _, item := range testCases {
		ret = oneEditAway(item[0].(string), item[1].(string))
		if ret != item[2].(bool) {
			t.Errorf("input:%v, output:%v", item, ret)
		}
	}
}
