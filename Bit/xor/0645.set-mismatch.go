package main

// https://leetcode.cn/problems/set-mismatch/

func findErrorNums(nums []int) []int {
	numsL := len(nums)

	// ---- 找出区别
	var numDiff int
	for num := 0; num < numsL; num++ { // hidden case 0
		numDiff ^= num
		numDiff ^= nums[num]
	}
	numDiff ^= numsL

	// TODO ----
	numDiff = numDiff & -numDiff // 切割 2个数字, 最低位1
	// 起始
	num1, num2 := 0, 0
	// 区间
	for n := 0; n < numsL; n++ {
		num := nums[n]
		if numDiff&num == 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}

		if numDiff&n == 0 {
			num1 ^= n
		} else {
			num2 ^= n
		}
	}

	// 结尾
	if numDiff&numsL == 0 {
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
