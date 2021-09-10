package main

import "math"

// https://leetcode-cn.com/problems/repeated-dna-sequences

func main() {
	findRepeatedDnaSequences("AAAAAAAAAAAAAAAAAAAAAAAA")
}

//func findRepeatedDnaSequences(s string) []string {
//	m := map[string]int{}
//	l1 := len(s)
//	l := 0
//	r := 10
//	for r <= l1 {
//		m[s[l:r]]++
//		l++
//		r++
//	}
//
//	var res []string
//	for i := range m {
//		if 1 < m[i] {
//			res = append(res, i)
//		}
//	}
//	return res
//}

func findRepeatedDnaSequences(s string) []string {
	L, l1 := 10, len(s)
	if l1 <= L {
		return nil
	}
	base := 4
	al := int(math.Pow(4, float64(L-1)))
	h := 0
	var nums = []int{}
	m1 := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	for i := 0; i < l1; i++ {
		nums = append(nums, m1[s[i]])
	}
	var res []string
	m2 := map[int]int{}
	for i := 0; i < L; i++ {
		h = h*base + nums[i]
	}
	m2[h]++
	for l := 1; l <= l1-L; l++ {
		h -= nums[l-1] * al // 减最左
		h *= base           // 提升
		h += nums[l+L-1]    // 加最右
		m2[h]++
		if m2[h] == 2 {
			res = append(res, s[l:l+L])
		}
	}
	return res
}
