package main

// https://leetcode.cn/problems/valid-perfect-square

func main() {
	println(isPerfectSquare(16))
}

// general
func isPerfectSquare(num int) bool {
	lo, hi := 0, num+1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		pow := mid * mid
		if pow < num {
			lo = mid + 1
		} else if num < pow {
			hi = mid // - 1 // TODO skip `equal` case 4 5
		} else if num == pow {
			return true
		}
	}
	return false // lo*lo == num
}

// offset
func isPerfectSquare(num int) bool {
	lo, hi := -1, num
	for lo < hi {
		mid := int(uint(lo+hi+1) >> 1)
		pow := mid * mid
		if pow < num {
			lo = mid // + 1 // TODO skip `equal` case 4 5
		} else if num < pow {
			hi = mid - 1
		} else if num == pow {
			return true
		}
	}
	return false // lo*lo == num
}

// equal
func isPerfectSquare(num int) bool {
	lo, hi := 0, num
	for lo <= hi {
		mid := int(uint(lo+hi) >> 1)
		pow := mid * mid
		if pow < num {
			lo = mid + 1
		} else if num < pow {
			hi = mid - 1
		} else if num == pow {
			return true
		}
	}
	return false
}
