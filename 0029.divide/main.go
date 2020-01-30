package main

import "math"

func divide(dividend int, divisor int) int {
	negative := dividend^divisor < 0
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	var multiple int32
	var aMultiple = divisor
	for divisor <= dividend {
		d := aMultiple
		m := int32(1)
		for divisor+d<<1 < dividend {
			d <<= 1
			m <<= 1
		}
		multiple += m

		divisor += d
	}
	if multiple < 0 { //处理溢出
		if negative {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}
	if negative {
		multiple = -multiple
	}
	if multiple > math.MaxInt32 {
		return math.MaxInt32
	}
	if multiple < math.MinInt32 {
		return math.MinInt32
	}
	return int(multiple)
}

/**
思路：移位倍增
 */
