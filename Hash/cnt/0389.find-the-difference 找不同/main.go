package main

// https://leetcode.cn/problems/find-the-difference

// ❓ 找出 x
// ⚠️ s = t + 随机byte(x)

func findTheDifferenceCnt(s string, t string) byte {
	// 统计
	chMpCnt := [26]int{}
	for i := range s {
		ch := s[i] - 'a'
		chMpCnt[ch]++
	}

	// map抵消
	for i := range t {
		ch := t[i] - 'a'
		if chMpCnt[ch] == 0 {
			return t[i]
		}
		chMpCnt[ch]--
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
