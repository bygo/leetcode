package main

// https://leetcode.cn/problems/find-all-anagrams-in-a-string

// ❓ 找到字符串中所有字母异位词(重排后相等)

func findAnagrams(s string, p string) []int {
	sL := len(s)
	pL := len(p)
	left := 0
	right := pL - 1
	if sL < pL {
		return nil
	}
	chMpCntP := [26]int{}
	for i := range p {
		ch := p[i] - 'a'
		chMpCntP[ch]++
	}

	// 计算相同长度的 vector
	chMpCntS := [26]int{}
	for left < right {
		ch := s[left] - 'a'
		chMpCntS[ch]++
		left++
	}

	var idxAnagrams []int
	left = 0
	for right < sL {
		ch := s[right] - 'a'
		chMpCntS[ch]++
		if chMpCntP == chMpCntS {
			idxAnagrams = append(idxAnagrams, left)
		}
		chMpCntS[s[left]-'a']--

		left++
		right++
	}

	return idxAnagrams
}
