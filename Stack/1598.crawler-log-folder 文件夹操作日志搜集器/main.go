package main

// https://leetcode.cn/problems/crawler-log-folder

func minOperations(logs []string) int {
	res := 0
	for i := range logs {
		switch logs[i][:2] {
		case "..": // -1
			if 0 < res {
				res--
			}
		case "./": // 忽略
		default:
			res++
		}
	}
	return res
}
