package main

/*
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

回文串不一定是字典当中的单词。



示例1：

输入："tactcoa"
输出：true（排列有"tacocat"、"atcocta"，等等）

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-permutation-lcci
*/

func canPermutePalindrome(s string) bool {
	record := make(map[rune]int, 0)

	for _, ch := range s {
		record[ch] ^= 1
	}

	// 最多一个字母出现基数次
	var flag bool
	for _, cnt := range record {
		if cnt == 1 {
			if flag {
				return false
			}
			flag = true
		}
	}

	return true
}
