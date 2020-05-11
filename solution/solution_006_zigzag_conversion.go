package solution

import (
	"bytes"
)

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
*/

// 0     6        12
// 1   5 7     11 13
// 2 4   8  10    14
// 3     9        15

// 数学归纳，k=0,1,2,3...
// 第0行: 0 6 12 ... 2*(numRows-1)*k
// 第1行：1 5 7 11 13 ... 2*(numRows-1)*k+1、2*(numRows-1)*(k+1)-1
// ...
// 第i行：2*(numRows-1)*k+i、2*(numRows-1)*(k+1)-i
// ...
// 第numRows-1行: 2*(numRows-1)*k+(numRows-1)
func convert(s string, numRows int) string {
	var (
		buffer         bytes.Buffer
		length         = len(s)
		lines          = numRows
		index1, index2 int
	)

	if length < numRows {
		lines = length
	}

	if lines <= 1 {
		return s
	}

	for i := 0; i < lines; i++ {
		if i == 0 {
			for k := 0; ; k++ {
				index1 = 2 * (lines - 1) * k
				if index1 >= length {
					break
				}
				buffer.WriteByte(s[index1])
			}
		} else if i == lines-1 {
			for k := 0; ; k++ {
				index1 = 2*(lines-1)*k + (lines - 1)
				if index1 >= length {
					break
				}
				buffer.WriteByte(s[index1])
			}
		} else {
			for k := 0; ; k++ {
				index1 = 2*(lines-1)*k + i
				if index1 >= length {
					break
				}
				buffer.WriteByte(s[index1])

				index2 = 2*(lines-1)*(k+1) - i
				if index2 >= length {
					break
				}
				buffer.WriteByte(s[index2])
			}
		}
	}

	return buffer.String()
}
