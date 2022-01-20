package main

import "math"

// https://leetcode-cn.com/problems/divide-two-integers

// ❓ 两数相除
// ⚠️ 随时判溢出

func quickAdd(x, y, z int32) bool {
	var numRes int32 = 0 // y*z
	for 0 < z {
		if 0 < z&1 {
			numRes += y
			if 0 <= numRes { // 溢出 z 过大
				return false
			}
			if numRes < x { //  y*z 超过x ，z 过大
				return false
			}
		}
		z >>= 1
		if 0 < z {
			// 下次倍增 y
			y <<= 1
			if 0 <= y { // 溢出 z 过大
				return false
			}
			if y < x { // y 自身超过 x ，z 过大
				return false
			}
		}
	}
	return true // z 在合理范围内
}

func divide(dividend, divisor int) int {
	// 负数范围更大
	minus := false

	var x, y = int32(dividend), int32(divisor)
	if 0 < dividend {
		x = -x
		minus = !minus
	}
	if 0 < divisor {
		y = -y
		minus = !minus
	}

	if x == 0 || y < x { // 被除数为 0
		return 0
	}

	if x == math.MinInt32 { // 被除数为最小值
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}

	//if y == math.MinInt32 { // 除数为最小值
	//	if x == math.MinInt32 {
	//		return 1
	//	}
	//	return 0
	//}

	var numRes int32 = 0
	var lo, hi int32 = 1, math.MaxInt32
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if quickAdd(x, y, mid) {
			numRes = mid
			if mid == math.MaxInt32 {
				break
				//return math.MaxInt32
			}
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	if minus {
		return -int(numRes)
	}
	return int(numRes)
}
