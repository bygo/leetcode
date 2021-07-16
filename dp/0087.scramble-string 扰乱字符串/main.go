package main

// https://leetcode-cn.com/problems/scramble-string

// 记忆递归
func isScramble(s1 string, s2 string) bool {
	// 记忆 减少重复计算
	var cache = make(map[string]bool)

	var dfs func(s1 string, s2 string) bool

	dfs = func(s1 string, s2 string) bool {
		var s = s1 + s2
		v, ok := cache[s]
		if ok {
			return v
		}

		if s1 == s2 {
			cache[s] = true
			return true
		}

		l := len(s1)

		var freq = map[byte]int{}
		for i := 0; i < l; i++ {
			freq[s1[i]]++
			freq[s2[i]]--
		}

		for _, v := range freq {
			if v != 0 {
				cache[s] = false
				return false
			}
		}

		//切割
		for i := 1; i < l; i++ {
			if dfs(s1[:i], s2[:i]) && dfs(s1[i:], s2[i:]) || dfs(s1[:i], s2[l-i:]) && dfs(s1[i:], s2[:l-i]) {
				cache[s] = true
				return true
			}
		}
		// 没有一个符合要求的
		cache[s] = false
		return false
	}
	return dfs(s1, s2)
}

// 三维
// f[i][j][l] 代表 字符串s1[i:i+l] s2[j:j+l]  是否和谐
func isScramble(s1 string, s2 string) bool {
	f := [30][30][31]bool{}
	l1 := len(s1)

	// 枚举长度 1 是否和谐
	for i := 0; i < l1; i++ {
		for j := 0; j < l1; j++ {
			f[i][j][1] = s1[i] == s2[j]
		}
	}

	// 枚举长度 2～l1 是否和谐
	for l := 2; l <= l1; l++ { // 长度变长
		// 枚举 A 中的起点位置
		for i := 0; i <= l1-l; i++ { //  5-2长度越长 字符串区间越小
			// 枚举 B 中的起点位置
			for j := 0; j <= l1-l; j++ { //  长度越长 字符串区间越小
				// 枚举划分位置
				for k := 1; k < l; k++ { // 1-l 长度区间 切割点
					if f[i][j][k] && f[i+k][j+k][l-k] || f[i][j+l-k][k] && f[i+k][j][l-k] {
						f[i][j][l] = true
						break
					}
				}
			}
		}
	}
	return f[0][0][l1]
}
