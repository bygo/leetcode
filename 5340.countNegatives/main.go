package main

func countNegatives(grid [][]int) int {
	var res int
	for _, nums := range grid {
		res += binary(nums)
	}
	return res
}

func binary(nums []int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= 0 {
			left++
		} else {
			if mid == 0 || nums[mid-1] >= 0 {
				return len(nums) - mid
			}
			right--
		}
	}
	return 0
}

/*
暴力法
 */
func countNegatives2(grid [][]int) int {
	var res int
	for _, nums := range grid {
		for k, v := range nums {
			if v < 0 {
				res += len(nums) - k
				break
			}
		}
	}
	return res
}
