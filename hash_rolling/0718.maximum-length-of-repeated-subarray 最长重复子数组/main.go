package main

// https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray

const (
	base = 131
)

func findLength(num1 []int, num2 []int) int {
	l1 := len(num1)
	l2 := len(num2)
	check := func(cnt int) bool {
		var h uint
		for i := 0; i < cnt; i++ {
			h = h*base + uint(num1[i])
		}
		b := map[uint]bool{h: true}
		mul := pow(base, cnt-1) // 最高位系数
		for i := 0; i < l1-cnt; i++ {
			h = h - uint(num1[i])*mul      // 减最高位 ，h可能小于 mod  变负数
			h = h*base + uint(num1[i+cnt]) // 剩下的数倍增
			b[h] = true
		}

		h = 0
		for i := 0; i < cnt; i++ {
			h = h*base + uint(num2[i])
		}
		if b[h] {
			return true
		}
		for i := 0; i < l2-cnt; i++ {
			h = h - uint(num2[i])*mul
			h = h*base + uint(num2[i+cnt])
			if b[h] {
				return true
			}
		}
		return false
	}

	left, right := 1, min(len(num1), len(num2))+1
	for left < right {
		mid := (left + right) >> 1
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

func pow(x, n int) uint {
	res := 1
	for n != 0 {
		if n&1 != 0 {
			res = res * x
		}
		x = x * x
		n >>= 1
	}
	return uint(res)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
