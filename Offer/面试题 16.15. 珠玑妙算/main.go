package main

// https://leetcode-cn.com/problems/master-mind-lcci/

func masterMind(solution string, guess string) []int {
	var res = make([]int, 2)
	m := map[byte]int{}
	for i := range solution {
		if solution[i] == guess[i] {
			res[0]++
		} else {
			m[solution[i]]++
			m[guess[i]]--
		}
	}
	var cnt int
	//for i := range m {
	//	if 0 < m[i] {
	//		cnt += m[i]
	//	}
	//}

	for i := range m {
		if m[i] < 0 {
			cnt -= m[i]
		}
	}
	res[1] = 4 - cnt - res[0]
	return res
}
