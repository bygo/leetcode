package main

import "strconv"

// https://leetcode.cn/problems/count-and-say/

func countAndSay(n int) string {
	var res = []byte("1$")
	var cur []byte
	var cnt = 1
	for ; 1 < n; n-- {
		l1 := len(res) - 1
		cur = []byte{}
		for i := 0; i < l1; i++ {
			if res[i] == res[i+1] {
				cnt++
			} else {
				cur = append(cur, strconv.Itoa(cnt)...)
				cur = append(cur, res[i])
				cnt = 1
			}
		}

		res = []byte{}
		res = append(res, cur...)
		res = append(res, '$')
	}
	return string(res[:len(res)-1])
}
