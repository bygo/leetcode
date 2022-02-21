package main

// https://leetcode-cn.com/problems/find-three-consecutive-integers-that-sum-to-a-given-number

// ❓ 找到和为给定整数的三个连续整数

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return nil
	}
	num /= 3
	return []int64{num - 1, num, num + 1}
}
