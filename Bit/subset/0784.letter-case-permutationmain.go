package main

// https://leetcode.cn/problems/letter-case-permutation

func letterCasePermutation(s string) []string {
	var subset int
	var buf = []byte(s)
	for i := range s {
		if s[i] <= '9' {
			continue
		}
		subset |= 1 << i
	}

	var strs []string
	sub := subset
	for {
		mask := sub
		for idx := range s {
			if mask>>idx&1 == 0 {
				buf[idx] = s[idx]
				continue
			}
			if s[idx] <= '9' {
				continue
			}
			if 'a' <= s[idx] && s[idx] <= 'z' {
				buf[idx] = s[idx] - 32
			} else {
				buf[idx] = s[idx] + 32
			}
		}
		strs = append(strs, string(buf))
		sub = (sub - 1) & subset
		if sub == subset {
			break
		}
	}
	return strs
}
