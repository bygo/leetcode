package main

// https://leetcode.cn/problems/maximum-product-of-word-lengths

// ❓ 最大单词长度乘积
// hash

func maxProduct(words []string) int {
	hashMpLen := map[int32]int{}
	var numMax int
	for _, word := range words {
		wordL := len(word)
		var hashCur int32
		for _, ch := range word {
			hashCur |= 1 << (ch - 'a')
		}

		// 相同hash 必须更大才计算
		if wordL <= hashMpLen[hashCur] {
			continue
		}
		hashMpLen[hashCur] = wordL
	}

	// 组合乘积
	for hashX, xL := range hashMpLen {
		for hashY, yL := range hashMpLen {
			if hashX&hashY != 0 {
				continue
			}
			numCur := xL * yL
			if numMax < numCur {
				numMax = numCur
			}
		}
	}
	return numMax
}
