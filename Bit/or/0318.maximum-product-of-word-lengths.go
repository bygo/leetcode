package main

// https://leetcode.cn/problems/maximum-product-of-word-lengths

// 2 <= words.length <= 1000
// 1 <= words[i].length <= 1000

func maxProduct(words []string) int {
	subMpLen := map[int32]int{}
	var numMax int
	for _, word := range words {
		wordL := len(word)
		var sub int32
		for _, ch := range word {
			sub |= 1 << (ch - 'a') // TODO 是否存在
		}

		if wordL <= subMpLen[sub] {
			continue
		}
		subMpLen[sub] = wordL
	}

	for subX, xL := range subMpLen {
		for subY, yL := range subMpLen {
			if subX&subY != 0 { // TODO
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
