package main

// https://leetcode-cn.com/problems/longest-duplicate-substring

const base = 1331

func longestDupSubstring(s string) string {
	l := len(s)
	var res string
	check := func(cnt int) bool {
		var h uint
		for i := 0; i < cnt; i++ {
			ch := uint(s[i])
			h = h*base + ch
		}

		b := map[uint]int{h: 1}
		mul := pow(base, cnt-1) // 最高位系数
		for i := 0; i < l-cnt; i++ {
			ch1 := uint(s[i])
			ch2 := uint(s[i+cnt])
			h = h - ch1*mul  // 减最高位 ，h可能小于 mod  变负数
			h = h*base + ch2 // 剩下的数倍增

			cur := s[i+1 : i+cnt+1]
			b[h]++
			if b[h] == 2 {
				res = cur
				return true
			}
		}
		return false
	}

	left, right := 1, l+1
	for left < right {
		mid := (left + right) >> 1
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return res
}

func pow(base, n int) uint {
	res := 1
	for n != 0 {
		if n&1 != 0 {
			res = res * base
		}
		base = base * base
		n >>= 1
	}
	return uint(res)
}
