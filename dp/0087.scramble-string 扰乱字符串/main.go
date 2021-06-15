package main

//Link: https://leetcode-cn.com/problems/scramble-string

// 递归
func isScramble3(s1 string, s2 string) bool {
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

		var d [26]int
		for i := 0; i < l; i++ {
			d[s1[i]-'a']++
			d[s2[i]-'a']--
		}

		for _, v := range d {
			if v != 0 {
				cache[s] = false
				return false
			}
		}

		//切割
		for i := 1; i < l; i++ {
			if dfs(s1[:i], s2[:i]) && dfs(s1[i:], s2[i:]) || dfs(s1[:i], s2[l-i:]) && dfs(s1[i:], s2[:l-i]) {
				cache[s] = true
				return cache[s]
			}
		}

		// 没有一个符合要求的
		cache[s] = false
		return false
	}
	return dfs(s1, s2)
}

func isScramble2(s1, s2 string) bool {
	n := len(s1)
	dp := make([][][]int8, n)
	for i := range dp {
		dp[i] = make([][]int8, n)
		for j := range dp[i] {
			dp[i][j] = make([]int8, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	// 第一个字符串从 i1 开始，第二个字符串从 i2 开始，子串的长度为 length
	// 和谐返回 1，不和谐返回 0
	var dfs func(i1, i2, l int) int8
	dfs = func(i1, i2, l int) int8 {
		var res int8
		d := &dp[i1][i2][l]
		if *d != -1 {
			return *d
		}
		defer func() { *d = res }()

		// 判断两个子串是否相等
		x, y := s1[i1:i1+l], s2[i2:i2+l]
		if x == y {
			return 1
		}

		// 判断是否存在字符 c 在两个子串中出现的次数不同
		freq := [26]int{}
		for i := range x {
			freq[x[i]-'a']++
			freq[y[i]-'a']--
		}
		for _, f := range freq {
			if f != 0 {
				return 0
			}
		}

		// 枚举分割位置
		for i := 1; i < l; i++ {
			// 不交换的情况
			if dfs(i1, i2, i) == 1 && dfs(i1+i, i2+i, l-i) == 1 {
				return 1
			}
			// 交换的情况
			if dfs(i1, i2+l-i, i) == 1 && dfs(i1+i, i2, l-i) == 1 {
				return 1
			}
		}
		return 0
	}
	return dfs(0, 0, n) == 1
}

// s1 = "great", s2 = "rgeat"
//

func main() {
	isScramble("great", "eatgr")
}

// dp[i][j][l] 代表 字符串s1[0:l] s2[0:l]  是否和谐
func isScramble(s1 string, s2 string) bool {
	dp := [30][30][31]bool{}
	n := len(s1)

	// 枚举长度 1 是否和谐
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	// 枚举长度 2～n 是否和谐
	for l := 2; l <= n; l++ { // 长度变长
		// 枚举 A 中的起点位置
		for i := 0; i <= n-l; i++ { //  5-2长度越长 字符串区间越小
			// 枚举 B 中的起点位置
			for j := 0; j <= n-l; j++ { //  长度越长 字符串区间越小
				// 枚举划分位置
				for k := 1; k < l; k++ { // 1-l 长度区间 切割点
					if dp[i][j][k] && dp[i+k][j+k][l-k] {
						dp[i][j][l] = true
						break
					}

					if dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][n]
}

func is(s1, s2 string) bool {
	dp := [30][30][31]bool{}
	n := len(s1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][1] = s1[i] == s2[j]
		}
	}

	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			for j := 0; j <= n-l; j++ {
				for k := 1; k < l; k++ {
					if dp[i][j][k] && dp[i+k][j+k][l-k] || dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][n]
}
