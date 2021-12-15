package main

import (
	"strconv"
)

// https://leetcode-cn.com/problems/subdomain-visit-count

// ❓ 子域名访问次数
// ⚠️ 10 www.by.com

func subdomainVisits(cpdomains []string) []string {
	var hostMpCnt = map[string]int{}
	for _, str := range cpdomains {
		// 出现次数
		var cnt int
		var idx int
		var strL = len(str)
		for str[idx] != ' ' {
			cnt = cnt*10 + int(str[idx]-'0')
			idx++
		}
		idx++

		// 域名
		for idx < strL {
			hostMpCnt[str[idx:]] += cnt
			for idx < strL && str[idx] != '.' {
				idx++
			}
			idx++
		}
	}
	var res []string
	for host, cnt := range hostMpCnt {
		res = append(res, strconv.Itoa(cnt)+" "+host)
	}
	return res
}
