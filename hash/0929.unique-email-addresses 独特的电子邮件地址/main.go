package main

import "strings"

// https://leetcode-cn.com/problems/unique-email-addresses

func numUniqueEmails(emails []string) int {
	m := map[string]struct{}{}
	for _, email := range emails {
		e := strings.Split(email, "@")
		cur := []byte{}
		for i := range e[0] {
			if e[0][i] == '.' {
				continue
			} else if e[0][i] == '+' {
				break
			} else {
				cur = append(cur, e[0][i])
			}
		}
		cur = append(cur, '@')
		cur = append(cur, e[1]...)
		m[string(cur)] = struct{}{}
	}
	return len(m)
}
