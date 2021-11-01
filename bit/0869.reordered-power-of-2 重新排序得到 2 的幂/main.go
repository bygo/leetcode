package main

import (
	"sort"
	"strconv"
)

// https://leetcode-cn.com/problems/reordered-power-of-2
func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0
}

func reorderedPowerOf2(n int) bool {
	nums := []byte(strconv.Itoa(n))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	l := len(nums)
	vis := make([]bool, l)
	var back func(dist, num int) bool
	var cnt int
	back = func(dist, num int) bool {
		if dist == l {
			return isPowerOfTwo(num)
		}
		for i, ch := range nums {
			cnt++
			if vis[i] || num == 0 && ch == '0' || 0 < i && ch == nums[i-1] && !vis[i-1] { //组合 位置相同 = 相同问题
				continue
			}
			vis[i] = true
			if back(dist+1, num*10+int(ch-'0')) {
				return true
			}
			vis[i] = false
		}
		return false
	}
	return back(0, 0)
}

func countDigits(n int) (cnt [10]int) {
	for 0 < n {
		cnt[n%10]++
		n /= 10
	}
	return
}

var powerOf2Digits = map[[10]int]bool{}

func init() {
	for n := 1; n <= 1e9; n <<= 1 {
		powerOf2Digits[countDigits(n)] = true
	}
}

func reorderedPowerOf2(n int) bool {
	return powerOf2Digits[countDigits(n)]
}
