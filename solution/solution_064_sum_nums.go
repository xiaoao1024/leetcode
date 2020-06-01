package solution

/*
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

示例 1：

输入: n = 3
输出: 6
示例 2：

输入: n = 9
输出: 45

限制：
1 <= n <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/qiu-12n-lcof
*/
func sumNums(n int) int {
	//if n == 0 {
	//	return 0
	//}

	//return n + sumNums(n-1)

	var (
		ans     int
		sumFunc func(int) bool
	)

	sumFunc = func(n int) bool {
		ans += n
		return n > 0 && sumFunc(n-1)
	}

	sumFunc(n)

	return ans
}
