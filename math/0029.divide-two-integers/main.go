package main

import "math"

// https://leetcode-cn.com/problems/divide-two-integers/

func divide(dividend int, divisor int) int {
	negative := dividend^divisor < 0
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	var sum int32
	var dd = divisor
	for dd <= dividend {
		d := divisor
		cur := int32(1)
		for dd+d<<1 < dividend {
			d <<= 1
			cur <<= 1
		}
		sum += cur
		dd += d
	}
	if sum < 0 {
		if negative {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}
	if negative {
		sum = -sum
	}
	if math.MaxInt32 < sum {
		return math.MaxInt32
	}
	if sum < math.MinInt32 {
		return math.MinInt32
	}
	return int(sum)
}

func quickAdd(x, y, z int) bool {
	res := 0
	add := y
	for 0 < z {
		if 0 < z&1 {
			if res < x-add {
				return false
			}
			res += add
		}
		if z != 1 {
			if add < x-add {
				return false
			}
			add += add
		}
		z >>= 1
	}
	return true
}

/**
思路：移位倍增
*/
