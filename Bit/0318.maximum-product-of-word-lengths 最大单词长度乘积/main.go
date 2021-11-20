package main

// https://leetcode-cn.com/problems/maximum-product-of-word-lengths

func maxProduct(words []string) int {
	l1 := len(words)
	m := map[int]int{}
	order := []int{}
	for i := 0; i < l1; i++ {
		l2 := len(words[i])
		cur := 0
		for j := 0; j < l2; j++ {
			cur |= 1 << (words[i][j] - 'a')
		}
		_, ok := m[cur]
		if !ok {
			order = append(order, cur)
		}
		if m[cur] < len(words[i]) {
			m[cur] = len(words[i])
		}
	}

	var res int
	l3 := len(order)
	for i := 0; i < l3; i++ {
		for j := i; j < l3; j++ {
			if order[i]&order[j] == 0 {
				cur := m[order[i]] * m[order[j]]
				if res < cur {
					res = cur
				}
			}
		}
	}
	return res
}
