package main

// https://leetcode.cn/problems/set-mismatch/

// ❓ 错误的集合

func findErrorNums(nums []int) []int {
	numsL := len(nums)

	// ⚠️ 找出区别
	var numStd int
	//numStd ^= nums[0]
	for n := 0; n < numsL; n++ {
		numStd ^= n
		numStd ^= nums[n]
	}
	numStd ^= numsL

	// ⚠️ 切割 2个数字, 最低位1
	numStd = numStd & -numStd

	// 起始
	num1, num2 := 0, 0
	// 区间
	for n := 0; n < numsL; n++ {
		num := nums[n]
		if numStd&num == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}

		if numStd&n == 0 {
			num1 ^= n
		} else {
			num2 ^= n
		}
	}

	// 结尾
	if numStd&numsL == 0 {
		num1 ^= numsL
	} else {
		num2 ^= numsL
	}

	for _, num := range nums {
		if num == num1 {
			return []int{num1, num2}
		}
	}
	return []int{num2, num1}
}
