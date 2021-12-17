package main

// https://leetcode-cn.com/problems/find-the-difference

// ❓ 找出 x
// ⚠️ s = t + 随机byte(x)

func findTheDifferenceCnt(s string, t string) byte {
	// 统计
	chMpIdx := [26]int{}
	for i := range s {
		ch := s[i] - 'a'
		chMpIdx[ch]++
	}

	// map抵消
	for i := range t {
		ch := t[i] - 'a'
		if chMpIdx[ch] == 0 {
			return ch + 'a'
		}
		chMpIdx[ch]--
	}
	return 0
}

func findTheDifferenceDistance(s string, t string) byte {
	// sum抵消
	var sum = 0
	for i := range s {
		sum = sum - int(s[i]) + int(t[i])
	}
	return byte(sum + int(t[len(t)-1]))
}

func findTheDifferenceXor(s string, t string) byte {
	// 异或抵消
	var sum byte
	for i := range s {
		sum ^= s[i] ^ t[i]
	}
	return sum ^ t[len(t)-1]
}
