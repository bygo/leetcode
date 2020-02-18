package main

import (
	"math"
)

func firstMissingPositive(nums []int) int {
	l := len(nums)
	contains := 0
	for i := 0; i < l; i++ { //检查是否包含1
		if nums[i] == 1 {
			contains++
			break
		}
	}

	if contains == 0 { //不包含1即答案
		return 1
	}
	if l == 1 { //包含且长度为1，  2即答案
		return 2
	}

	for i := 0; i < l; i++ { //把所有大数,负数，全部转换为1，因为值必定 res <= l+1
		if nums[i] <= 0 || l < nums[i] {
			nums[i] = 1
		}
	}

	for i := 0; i < l; i++ { //以符号的正负 当作hash表
		index := int(math.Abs(float64(nums[i]))) //获取需要改变的索引
		if index == l {
			nums[0] = - int(math.Abs(float64(nums[0])))
		} else {
			nums[index] = - int(math.Abs(float64(nums[index])))
		}
	}

	for i := 1; i < l; i++ { //不存在 即答案
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 { // 不存在 即答案
		return l
	}
	return l + 1 //都存在，后序l+1为答案
}

func main() {
	firstMissingPositive([]int{-2, -1, 1, 2, 3, 4})
}
