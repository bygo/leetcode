package main

import "leetcode"

/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

 */
func reverse(x int) int {
	var y int
	for x != 0 {
		y = y*10 + x%10
		if y < -1 << 31 || y > 1<<31-1 {
			return 0
		}
		x /= 10
	}
	return y
}

func main() {
	leetcode.D(func() interface{} {
		return reverse(321)
	})
}

/**
思路：
1.不断对 x 取余（取尾数），除10（舍尾数）
2.上次结果 y 尾数补0
 */
