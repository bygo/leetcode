package main

import "math"

// https://leetcode-cn.com/problems/divide-two-integers

// ❓ 两数相除
// ⚠️ 随时判溢出

// 快速乘 divisor * quotient
func quickAdd(dividend int32, divisor, quotient int32) bool {
	var numRes int32
	for 0 < quotient {
		if quotient&1 == 1 {
			numRes += divisor
			// 溢出
			if 0 <= numRes || numRes < dividend {
				return false
			}
		}

		quotient >>= 1
		if 0 < quotient {
			divisor <<= 1
			// 溢出
			if 0 <= divisor || divisor < dividend {
				return false
			}
		}

	}
	return true
}

func divide(x, y int) int {
	// 负数范围更大
	minus := x^y < 0

	var dividend, divisor = int32(x), int32(y)
	if 0 < x {
		dividend = -dividend
	}
	if 0 < y {
		divisor = -divisor
	}

	// 特殊用例，被除数为 0 | 除数 有效位更大
	if dividend == 0 || divisor < dividend {
		return 0
	}

	// 特殊用例，被除数为最小值 超过 math.MaxInt32 的精度
	if dividend == math.MinInt32 {
		if y == 1 {
			return math.MinInt32
		}
		if y == -1 {
			return math.MaxInt32
		}
	}
	// 除数为最小值，可以精准计算，无需判断
	//if divisor == math.MinInt32 {
	//	if dividend == math.MinInt32 {
	//		return 1
	//	}
	//	return 0
	//}

	var numRes int32 = 0
	var lo, hi int32 = 1, math.MaxInt32
	// `<=` 充分利用整数范围
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		// 合法
		if quickAdd(dividend, divisor, mid) {
			numRes = mid
			if mid == math.MaxInt32 {
				break
				//return math.MaxInt32
			}
			lo = mid + 1
		} else {
			// 溢出 (mid 过大)
			hi = mid - 1
		}
	}

	if minus {
		// 还原
		return -int(numRes)
	}
	return int(numRes)
}
