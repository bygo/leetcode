package main

/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
*/

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	var reverse int
	for x > reverse {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return x == reverse || x == reverse/10
}

/**
思路：
1.获取长度为 length/2 +1 的反转数字 ，对比两个数字大小
假设12321=>12  123    12=123/10
假设1221=>12   12 	 12=12
*/
