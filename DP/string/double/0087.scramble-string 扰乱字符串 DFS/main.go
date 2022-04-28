package main

// https://leetcode-cn.com/problems/scramble-string

// 记忆递归
// 旋转后相等

func isScramble(s1 string, s2 string) bool {
	// 记忆 减少重复计算
	var cache = make(map[string]bool)

	var dfs func(s1 string, s2 string) bool

	dfs = func(s1 string, s2 string) bool {
		var k = s1 + s2
		b, ok := cache[k]
		if ok {
			return b
		}

		l := len(s1)

		// 直接相等
		if s1 == s2 {
			cache[k] = true
			return true
		}
		// 如果不相等，就开始旋转

		// 频数优化
		// 字符数不同，无论旋转，都不可能相等
		var chMpCnt = map[byte]int{}
		for i := 0; i < l; i++ {
			chMpCnt[s1[i]]++
			chMpCnt[s2[i]]--
		}
		for _, cnt := range chMpCnt {
			if cnt != 0 {
				cache[k] = false
				return false
			}
		}

		// 依次切割后旋转
		cache[k] = false
		for i := 1; i < l; i++ {
			if dfs(s1[:i], s2[:i]) && dfs(s1[i:], s2[i:]) || dfs(s1[:i], s2[l-i:]) && dfs(s1[i:], s2[:l-i]) {
				cache[k] = true
				break
			}
		}
		return cache[k]
	}
	return dfs(s1, s2)
}
