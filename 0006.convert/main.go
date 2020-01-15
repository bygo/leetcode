package main

import "leetcode"

/**
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

 */

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	stackMap := make([][]byte, numRows)
	i, j := 0, 1
	for _, v := range s {
		stackMap[i] = append(stackMap[i], byte(v))
		i += j
		if i == 0 || i == numRows-1 {
			j -= j * 2
		}
	}
	var z []byte
	for _, stack := range stackMap {
		z = append(z, stack...)
	}
	return string(z)
}

func main() {
	leetcode.D(func() interface{} {
		return convert("LEETCODEISHIRING", 4)
	})
}

/**
思路1：
1. 设置 1-numRows 切片收集字符,i 为当前行数,j 为步长
2. 步长 j 为 -1 时 向上退一行，为 1 时向下进一行，遇到边界取反

思路2
1. Z 型字符的各个位置可以以方程式计算得出
numRows=n=4
k=key/(numRows+2)
0     6      12        k(2*n-2)							0,6,12...
1  5  7  11  13 17     k(2*n-2)+1 (k+1)(2*n-2)-1		[1,5] [7,11] [13,1]...
2  4  8  10  14 16     k(2*n-2)+2 (k+1)(2*n-2)-2		[2,4] [8,10] [14,16]...
3     9      15        k(2*n-2)+3						3,9,15...
 */
