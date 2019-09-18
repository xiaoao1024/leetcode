package solution

/*
统计所有小于非负整数 n 的质数的数量。

示例:

输入: 10
输出: 4
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
*/

// 厄拉多塞筛法
func countPrimes(n int) int {
	isNotPrimes := make([]bool, n)

	var cnt int
	for i := 2; i < n; i++ {
		if !isNotPrimes[i] {
			cnt++
			for j := 2; i*j < n; j++ {
				isNotPrimes[i*j] = true
			}
		}
	}

	return cnt
}
