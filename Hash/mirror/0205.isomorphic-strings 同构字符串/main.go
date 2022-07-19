package main

// https://leetcode.cn/problems/isomorphic-strings

// ❓ 同构字符串

func isIsomorphic(s string, t string) bool {
	// 双向绑定
	ch1MpCh2 := map[byte]byte{}
	ch2MpCh1 := map[byte]byte{}
	sL := len(s)
	for i := 0; i < sL; i++ {
		ch1 := s[i]
		ch2 := t[i]

		if ch1MpCh2[ch1] != 0 && ch1MpCh2[ch1] != ch2 || ch2MpCh1[ch2] != 0 && ch2MpCh1[ch2] != ch1 {
			return false
		}
		ch1MpCh2[ch1] = ch2
		ch2MpCh1[ch2] = ch1
	}

	return true
}
