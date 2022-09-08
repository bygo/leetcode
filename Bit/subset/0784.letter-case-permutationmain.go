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
			//if s[idx] <= '9' {
			//	continue
			//}
			buf[idx] = s[idx] ^ byte(mask>>idx&1<<5) // TODO 大小写互转，数字忽略
		}
		strs = append(strs, string(buf))
		if sub == 0 {
			break
		}
		sub = (sub - 1) & subset
		//if sub == subset {
		//	break
		//}
	}
	return strs
}
