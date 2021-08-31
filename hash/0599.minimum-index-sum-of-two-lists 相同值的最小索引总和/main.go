package main

// https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists

func findRestaurant(list1 []string, list2 []string) []string {
	m := map[string]int{}
	for i := range list1 {
		m[list1[i]] = i
	}

	var res []string
	var min = 1<<63 - 1
	for i := range list2 {
		index, ok := m[list2[i]]
		if ok {
			cur := index + i
			if cur < min {
				res = []string{list2[i]}
				min = cur
			} else if cur == min {
				res = append(res, list2[i])
			}
		}
	}
	return res
}
