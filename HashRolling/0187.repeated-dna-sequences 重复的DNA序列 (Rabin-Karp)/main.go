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

var bases = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

const cnt = 10
const base = 4

func findRepeatedDnaSequences(s string) []string {
	l1 := len(s)
	if l1 <= cnt {
		return nil
	}
	mul := int(math.Pow(4, float64(cnt-1))) // 最高位
	h := 0
	m := map[int]int{}
	for i := 0; i < cnt; i++ {
		h = h*base + bases[s[i]]
	}
	m[h]++

	hash := func(l, r int) int {
		h = h - bases[s[l]]*mul // 减最左
		h *= base               // 提升
		h += bases[s[r]]        // 加最右
		return h
	}

	var res []string
	for l := 0; l < l1-cnt; l++ {
		h = hash(l, l+cnt)
		m[h]++
		if m[h] == 2 {
			res = append(res, s[l+1:l+cnt+1])
		}
	}
	return res
}
