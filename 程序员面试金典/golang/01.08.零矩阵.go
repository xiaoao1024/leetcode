package main

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
func setZeroes(matrix [][]int) {
	var (
		rowLen = len(matrix)
		colLen = len(matrix[0])

		recordRow = make([]int, colLen)
		recordCol = make([]int, rowLen)
	)
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if matrix[i][j] == 0 {
				recordRow[j] = 1
				recordCol[i] = 1
			}
		}
	}

	// 检查第一列
	for i := 0; i < rowLen; i++ {
		if recordCol[i] == 1 {
			for j := 0; j < colLen; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// 检查第一行
	for j := 0; j < colLen; j++ {
		if recordRow[j] == 1 {
			for i := 0; i < rowLen; i++ {
				matrix[i][j] = 0
			}
		}
	}
}
