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
	var multiple int
	var aMultiple = divisor
	for divisor <= dividend {
		d := aMultiple
		m := 1
		for divisor+d<<1 < dividend {
			d <<= 1
			m <<= 1
		}
		multiple += m
		divisor += d
	}
	if negative {
		multiple = -multiple
	}
	if multiple > math.MaxInt32 {
		multiple = math.MaxInt32
	}
	if multiple < math.MinInt32 {
		multiple = math.MinInt32
	}
	return multiple
}

/**
思路：移位倍增
 */
