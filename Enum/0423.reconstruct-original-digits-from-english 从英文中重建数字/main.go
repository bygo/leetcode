package main

// https://leetcode.cn/problems/reconstruct-original-digits-from-english

func originalDigits(s string) string {
	m1 := [26]int{}

	for i := range s {
		m1[s[i]-'a']++
	}
	m2 := [10]int{}
	m2[0] = m1['z'-'a']
	m2[2] = m1['w'-'a']
	m2[4] = m1['u'-'a']
	m2[8] = m1['x'-'a']

	m2[3] = m1['h'-'a'] - m2[8]
	m2[5] = m1['f'-'a'] - m2[4]
	m2[7] = m1['s'-'a'] - m2[6]

	m2[9] = m1['i'-'a'] - m2[5] - m2[6] - m2[8]
	m2[1] = m1['n'-'a'] - m2[7] - 2*m2[9]

	var res []byte
	for i := 0; i < 10; i++ {
		for j := 0; j < m2[i]; j++ {
			res = append(res, byte(i+'0'))
		}
	}
	return string(res)
}
