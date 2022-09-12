package main

// https://leetcode.cn/problems/letter-case-permutation

func letterCasePermutation(s string) []string {
	var subset int
	var buf = []byte(s)
	var idxes []int
	for i := range s {
		if 'A' <= s[i] {
			idxes = append(idxes, i)
			subset |= 1 << i
		}
	}

	var strs []string
	sub := subset
	for {
		for _, idx := range idxes {
			buf[idx] = s[idx] ^ byte(sub>>idx&1<<5)
		}
		strs = append(strs, string(buf))
		if sub == 0 {
			break
		}
		sub = (sub - 1) & subset
	}
	return strs
}
