package main

// https://leetcode-cn.com/problems/maximum-product-of-word-lengths

// ❓ 最大单词长度乘积

func maxProduct(words []string) int {
	hashMpLen := map[int32]int{}
	var numMax int
	for _, word := range words {
		wordL := len(word)
		var hashCur int32 = 0
		for j := 0; j < wordL; j++ {
			hashCur |= 1 << (word[j] - 'a')
		}

		// 组合乘积
		if hashMpLen[hashCur] < wordL {
			for hash, length := range hashMpLen {
				if hash&hashCur == 0 {
					numCur := length * wordL
					if numMax < numCur {
						numMax = numCur
					}
				}
			}
			hashMpLen[hashCur] = wordL
		}
	}
	return numMax
}
