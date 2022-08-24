package main

// https://leetcode.cn/problems/maximum-product-of-word-lengths

// 2 <= words.length <= 1000
// 1 <= words[i].length <= 1000

func maxProduct(words []string) int {
	hashMpLen := map[int32]int{}
	var numMax int
	for _, word := range words {
		wordL := len(word)
		var hashCur int32
		for _, ch := range word {
			hashCur |= 1 << (ch - 'a')
		}

		if wordL <= hashMpLen[hashCur] {
			continue
		}
		hashMpLen[hashCur] = wordL
	}

	for hashX, xL := range hashMpLen {
		for hashY, yL := range hashMpLen {
			if hashX&hashY != 0 { // TODO
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
