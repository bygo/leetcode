package main

// https://leetcode-cn.com/problems/find-the-difference

// ❓ 找出 x
// ⚠️ s = t + 随机byte(x)

func findTheDifferenceXor(s string, t string) byte {
	// 异或抵消
	var sum byte
	for i := range s {
		sum ^= s[i] ^ t[i]
	}
	return sum ^ t[len(t)-1]
}
