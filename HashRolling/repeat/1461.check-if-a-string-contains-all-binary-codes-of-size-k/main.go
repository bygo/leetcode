package main

// https://leetcode-cn.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k

// ❓ 检查一个字符串是否包含所有长度为 K 的二进制子串

func hasAllCodes(s string, k int) bool {
	sL := len(s)
	if sL < k {
		return false
	}
	cntNum := 1 << k
	numMp := make([]bool, cntNum)
	var cntBit int

	var num int
	for i := 0; i < k; i++ {
		num <<= 1
		num += int(s[i] - '0')
	}

	numMp[num] = true
	cntBit++
	numStand := 1<<k - 1
	for i := k; i < sL; i++ {
		num <<= 1
		num += int(s[i] - '0')
		num &= numStand
		if !numMp[num] {
			numMp[num] = true
			cntBit++
		}
	}
	return cntBit == cntNum
}
