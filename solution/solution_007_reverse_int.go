package solution

import (
	"math"
)

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例1：
 输入：123
 输出：321

示例2：
 输入：-123
 输出：-321

示例3：
 输入：120
 输出：21

注意: 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

var (
	intMax = int(math.Pow(2, 31)) - 1    // 2147483647
	intMin = (-1) * int(math.Pow(2, 31)) // -2147483648
)

func reverse(x int) int {
	var (
		q   = x
		r   int
		res int
	)

	for q != 0 {
		r = q % 10

		if res > intMax/10 || (res == intMax/10 && r > 7) {
			return 0
		} else if res < intMin/10 || (res == intMin/10 && r < -8) {
			return 0
		}

		res = res*10 + r
		q = q / 10
	}
	return res
}
