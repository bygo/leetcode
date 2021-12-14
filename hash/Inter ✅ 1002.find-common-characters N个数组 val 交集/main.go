package main

// https://leetcode-cn.com/problems/find-common-characters

func commonChars(words []string) {
	var res []string
	minFreq := [26]int{}
	for i := range minFreq {
		minFreq[i] = 1<<63 - 1
	}
	for _, word := range words {
		freq := [26]int{}
		for _, b := range word {
			freq[b-'a']++
		}
		for i, f := range freq[:] {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}

	// 字典序
	for i := 0; i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			res = append(res, string(rune(i+'a')))
		}
	}
	return
}
