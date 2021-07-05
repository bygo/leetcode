package main

// Link: https://leetcode-cn.com/problems/destination-city

func destCity(paths [][]string) string {
	m := map[string]struct{}{}
	for _, path := range paths {
		m[path[0]] = struct{}{}
	}
	for _, path := range paths {
		_, ok := m[path[1]]
		if !ok {
			return path[1]
		}
	}
	return ""
}
