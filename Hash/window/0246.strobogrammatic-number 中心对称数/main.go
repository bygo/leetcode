package main

// https://leetcode-cn.com/problems/strobogrammatic-number

// ❓ 中心对称数

func isStrobogrammatic(num string) bool {
	chMpCh := map[byte]byte{'6': '9', '9': '6', '8': '8', '0': '0', '1': '1'}

	left, right := 0, len(num)-1
	for left <= right {
		chLeft := num[left]
		chRight := num[right]
		if chMpCh[chLeft] != chRight || chMpCh[chLeft] == 0 {
			return false
		}
		left++
		right--
	}
	return true
}
