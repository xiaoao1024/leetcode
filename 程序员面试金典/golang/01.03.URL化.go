package main

/*
URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）



示例 1：

输入："Mr John Smith    ", 13
输出："Mr%20John%20Smith"
示例 2：

输入："               ", 5
输出："%20%20%20%20%20"


提示：

字符串长度在 [0, 500000] 范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/string-to-url-lcci
*/

func replaceSpaces(S string, length int) string {
	var (
		arr  = []rune(S)
		i, j int
	)

	// 计算空格数量
	spaceCnt := countSpace(arr, length)
	// 新字符串长度
	len2 := length + 2*spaceCnt
	// i指向原字符串最后位置
	i = length - 1
	// j指向新字符串最后位置
	j = len2 - 1

	for ; i >= 0; i-- {
		if arr[i] == ' ' {
			arr[j] = '0'
			arr[j-1] = '2'
			arr[j-2] = '%'
			j -= 3
		} else {
			arr[j] = arr[i]
			j--
		}
	}

	return string(arr[:len2])
}

func countSpace(arr []rune, length int) int {
	var count int
	for i := 0; i < length; i++ {
		if arr[i] == ' ' {
			count++
		}
	}

	return count
}
