package main

// https://leetcode-cn.com/problems/zigzag-conversion/

// 1:08
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var cur, step = 0, 1
	var stack = make([][]byte, numRows)
	for i := range s {
		stack[cur] = append(stack[cur], s[i])
		cur += step
		if cur == 0 || cur == numRows-1 {
			step = -step
		}
	}

	var res []byte
	for _, b := range stack {
		res = append(res, b...)
	}

	return string(res)
}

/**
Z 型字符的各个位置也可以以方程式计算得出
numRows=n=4
k=key/(numRows+2)
0     6      12        k(2*n-2)							0,6,12...
1  5  7  11  13 17     k(2*n-2)+1 (k+1)(2*n-2)-1		[1,5] [7,11] [13,1]...
2  4  8  10  14 16     k(2*n-2)+2 (k+1)(2*n-2)-2		[2,4] [8,10] [14,16]...
3     9      15        k(2*n-2)+3						3,9,15...
*/
