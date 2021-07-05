package main

import (
	"bytes"
	"strconv"
)

// Link: https://leetcode-cn.com/problems/subdomain-visit-count

func subdomainVisits(cpdomains []string) []string {
	var m = map[string]int{}

	for _, c := range cpdomains {
		var num int
		var i int
		var n = len(c)
		for c[i] != ' ' {
			num = num*10 + int(c[i]-'0')
			i++
		}
		i += 1

		var words [][]byte
		cur := []byte{}
		for i < n {
			if c[i] == '.' {
				words = append(words, cur)
				cur = []byte{}
			} else {
				cur = append(cur, c[i])
			}
			i++
		}
		if 0 < len(cur) {
			words = append(words, cur)
		}

		l := len(words)
		for j := 0; j < l; j++ {
			w := bytes.Join(words[j:], []byte{'.'})
			m[string(w)] += num
		}
	}
	var res []string
	for k, v := range m {
		num := strconv.Itoa(v)
		res = append(res, num+" "+k)
	}
	return res
}
