package main

// https://leetcode-cn.com/problems/find-original-array-from-doubled-array

func findOriginalArray(changed []int) []int {
	l := len(changed)
	if l%2 == 1 {
		return nil
	}
	m := [100001]int{}
	var cnt int
	for i := range changed {
		m[changed[i]]++
		cnt++
	}

	var res []int
	for i := 0; i <= 50000; i++ {
		// 必须从小开始抵消
		for 0 < m[i] {
			m[i]--
			m[i*2]--
			cnt -= 2
			res = append(res, i)
			if m[i*2] < 0 {
				return nil
			}
		}
	}
	if 0 < cnt {
		return nil
	}
	return res
}
