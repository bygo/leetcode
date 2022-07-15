package main

// https://leetcode-cn.com/problems/sqrtx

func mySqrt(x int) int {
	//if x == 0 || x == 1 {
	//	return x
	//}
	lo, hi := 0, x+1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		cur := mid * mid
		if cur < x {
			lo = mid + 1
		} else if x < cur {
			hi = mid
		} else if x == cur {
			return mid
		}
	}
	return lo - 1
}

func mySqrt(x int) int {
	var lo = 0
	var hi = x
	for lo <= hi {
		mid := int(uint(lo+hi) >> 1)
		cur := mid * mid
		if cur < x {
			lo = mid + 1
		} else if x < cur {
			hi = mid - 1
		} else if x == cur {
			return mid
		}
	}
	return lo - 1
}
