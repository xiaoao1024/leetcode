package solution

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
*/

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	var (
		start, end, maxLen int
		start1, end1, len1 int
		start2, end2, len2 int
	)

	for i := 0; i < len(s); i++ {
		start1, end1 = indexLongestPalindromeSubstring(s, i, i)
		start2, end2 = indexLongestPalindromeSubstring(s, i, i+1)

		len1 = end1 - start1 + 1
		len2 = end2 - start2 + 1

		if len1 > len2 {
			if len1 > maxLen {
				start, end, maxLen = start1, end1, len1
			}
		} else {
			if len2 > maxLen {
				start, end, maxLen = start2, end2, len2
			}
		}
	}

	return s[start : end+1]
}

// 以[left, right]为中心的最长回文子串[start, end]
func indexLongestPalindromeSubstring(s string, left int, right int) (int, int) {
	var (
		length     = len(s)
		start, end int
	)

	if left > right || right >= length || left < 0 {
		return 0, 0
	}

	for left >= 0 && right < length && s[left] == s[right] {
		start, end = left, right
		left--
		right++
	}

	return start, end
}
