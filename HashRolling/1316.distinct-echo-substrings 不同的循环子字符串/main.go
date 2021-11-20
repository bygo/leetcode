package main

// https://leetcode-cn.com/problems/distinct-echo-substrings

const base uint = 131

func distinctEchoSubstrings(text string) int {
	n := len(text)
	pre := make([]uint, n+1)
	mul := make([]uint, n+1)
	mul[0] = 1
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1]*base + uint(text[i-1])
		mul[i] = mul[i-1] * base
	}

	hash := func(l, r int) uint {
		return pre[r+1] - pre[l]*mul[r-l+1] // 减去多乘  等于 0 ~ r-l 阶
	}
	var res int
	m := map[uint]bool{}

	for l := 0; l < n; l++ { // 第一串起始位
		r := l + 1 // 第二串起始位
		for {
			cnt := r - l - 1 // 距离
			if n <= r+cnt {
				break
			}
			left := hash(l, l+cnt) // 前置hash
			if !m[left] && left == hash(r, r+cnt) {
				res++
				m[left] = true
			}
			r++
		}
	}
	return res
}
