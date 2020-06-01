package solution

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, s := range strs {
		fmt.Println(prefix)
		prefix = lcp(prefix, s)
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func lcp(s1, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}

	var (
		prefix = make([]byte, 0)
	)

	for i := 0; i < len(s1); i++ {
		if i < len(s2) && s1[i] == s2[i] {
			prefix = append(prefix, s1[i])
		} else {
			break
		}
	}

	return string(prefix)
}
