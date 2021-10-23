package main

// https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray

func findLength(num1 []int, num2 []int) int {
	l1 := len(num1)
	l2 := len(num2)
	var maxLength func(a, b, l int) int
	maxLength = func(a, b, l int) int {
		var cur, k int
		for i := 0; i < l; i++ {
			if num1[a+i] == num2[b+i] {
				k++
			} else {
				k = 0
			}
			cur = max(cur, k)
		}
		return cur
	}
	var res int
	for i := 0; i < l1; i++ {
		l := min(l2, l1-i)
		res = max(res, maxLength(i, 0, l))
	}

	for i := 0; i < l2; i++ {
		l := min(l1, l2-i)
		res = max(res, maxLength(0, i, l))
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
