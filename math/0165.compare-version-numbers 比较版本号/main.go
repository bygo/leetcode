package main

// https://leetcode.cn/problems/compare-version-numbers

func compareVersion(version1 string, version2 string) int {
	l1 := len(version1)
	l2 := len(version2)
	var i, j int
	for i < l1 || j < l2 {
		var x, y int
		for i < l1 && version1[i] != '.' {
			x = x*10 + int(version1[i]-'0')
			i++
		}
		i++

		for j < l2 && version2[j] != '.' {
			y = y*10 + int(version2[j]-'0')
			j++
		}
		j++
		if x < y {
			return 1
		} else if y < x {
			return -1
		}
	}
	return 0
}
