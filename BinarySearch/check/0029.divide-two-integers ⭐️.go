package main

import "math"

// https://leetcode.cn/problems/divide-two-integers

func quickAdd(dividend int32, divisor int32, quotient int32) bool {
	var numRes int32
	for 0 < quotient {
		if quotient&1 == 1 {
			numRes += divisor
			if 0 <= numRes || numRes < dividend {
				// `numRes` overflow
				return false
			}
		}

		quotient >>= 1
		if 0 < quotient {
			divisor <<= 1
			if 0 <= divisor || divisor < dividend {
				// `divisor` overflow
				return false
			}
		}
	}
	return true
}

func divide(x int, y int) int {
	// Symbol bits repel each other
	minus := x^y < 0

	var dividend, divisor = int32(x), int32(y)
	// Negative numbers have a larger overflow range
	if 0 < x {
		dividend = -dividend
	}
	if 0 < y {
		divisor = -divisor
	}

	// Special case: the dividend is 0
	if dividend == 0 {
		return 0
	}

	// Special case: divisor larger
	if divisor < dividend {
		return 0
	}

	// Special case: The dividend is the `math.MaxInt32`
	if dividend == math.MinInt32 {
		if y == 1 {
			return math.MinInt32
		}
		if y == -1 {
			return math.MaxInt32
		}
	}
	// The divisor is the `math.MinInt32`
	// can be calculated accurately

	//if divisor == math.MinInt32 {
	//	if dividend == math.MinInt32 {
	//		return 1
	//	}
	//	return 0
	//}

	var numRes int32 = 0
	var lo, hi int32 = 1, math.MaxInt32
	// `<=` Take advantage of integer ranges
	for lo <= hi {
		mid := int32(uint32(lo+hi) >> 1)
		if quickAdd(dividend, divisor, mid) {
			numRes = mid
			if numRes == math.MaxInt32 {
				break
			}
			// `math.MaxInt32 + 1 == -2147483648 `will cause `lo<=hi` Infinite loop
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

func main() {
	var i int32 = math.MaxInt32
	println(i + 1)
}
