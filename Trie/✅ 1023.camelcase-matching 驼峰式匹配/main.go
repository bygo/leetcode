package main

// https://leetcode.cn/problems/camelcase-matching

// ❓ 驼峰式匹配

func camelMatch(queries []string, pattern string) []bool {
	bools := make([]bool, len(queries))
	patternL := len(pattern)

	fn := func(query string) bool {
		idxP := 0
		for j := range query {
			ch := query[j]
			if idxP < patternL && ch == pattern[idxP] {
				idxP++
			} else if 'A' <= ch && ch <= 'Z' {
				// 没有匹配到的 其他大写
				return false
			} else if 'a' <= ch && ch <= 'z' {
				// 忽略
				continue
			}
		}
		// 匹配结束
		return idxP == patternL
	}

	for i, query := range queries {
		bools[i] = fn(query)
	}
	return bools
}
