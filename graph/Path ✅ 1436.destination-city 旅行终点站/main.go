package main

// https://leetcode.cn/problems/destination-city

func destCity(paths [][]string) string {
	m := map[string]struct{}{}
	for _, path := range paths {
		m[path[0]] = struct{}{}
	}
	for _, path := range paths {
		_, ok := m[path[1]] // 终点站没有下一站
		if !ok {
			return path[1]
		}
	}
	return ""
}
