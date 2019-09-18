package solution

/*
编写一个算法来判断一个数是不是“快乐数”。

一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。

示例:

输入: 19
输出: true
解释:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1

*/

// 问题分析：
// 19 -> 82 -> 68 -> 100 -> 1, 在每一步计算之后，如果出现1，说明这是一个快乐数，但是如果没有出现1，而是出现某种循环，那么这个计算过程就会一直持续下去，因此需要判断“链表有环”
// “链表有环”问题可以用快慢指针
func isHappy(n int) bool {
	low := next(n)
	fast := next(next(n))

	for low != fast {
		low = next(low)
		fast = next(next(fast))
	}

	return fast == 1
}

func next(num int) int {
	var sum, factor int

	for num > 0 {
		factor = num % 10
		sum += factor * factor
		num /= 10
	}

	return sum
}
