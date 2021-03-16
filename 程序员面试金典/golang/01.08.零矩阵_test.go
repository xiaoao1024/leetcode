package main

import (
	"fmt"
	"testing"
)

/*
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。

示例 1：

输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2：

输入：
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出：
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]

*/

// 思路：用第一行和第一列来标记清理的行与列
func Test_setZeroes(t *testing.T) {
	m1 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	fmt.Println(m1)
	setZeroes(m1)
	fmt.Println(m1)

	m2 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	fmt.Println(m2)
	setZeroes(m2)
	fmt.Println(m2)
}
