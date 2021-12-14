package main

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/subdomain-visit-count

// 10 www.by.com
func subdomainVisits(cpdomains []string) []string {
	var m = map[string]int{}
	var cur []byte

	for _, c := range cpdomains {
		// 出现次数
		var cnt int
		var idx int
		var n = len(c)
		for c[idx] != ' ' {
			cnt = cnt*10 + int(c[idx]-'0')
			idx++
		}
		idx += 1

		// 域名全称
		var words []string
		for idx < n {
			if c[idx] == '.' {
				words = append(words, string(cur))
				cur = []byte{}
			} else {
				cur = append(cur, c[idx])
			}
			idx++
		}
		if 0 < len(cur) {
			words = append(words, string(cur))
			cur = []byte{}
		}

		// 计算域名所有子域
		// www.by.com
		// by.com
		// com
		l := len(words)
		for j := 0; j < l; j++ {
			w := strings.Join(words[j:], ".")
			m[w] += cnt
		}
	}
	var res []string
	for k, v := range m {
		res = append(res, strconv.Itoa(v)+" "+k)
	}
	return res
}
