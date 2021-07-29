package main

// https://leetcode-cn.com/problems/longest-common-prefix/

// 1:00
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var l1 = len(strs)
	var l2 = len(strs[0])
	for i := 0; i < l2; i++ {
		for j := 1; j < l1; j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
