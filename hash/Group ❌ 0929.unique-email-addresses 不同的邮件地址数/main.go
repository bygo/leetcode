package main

// https://leetcode-cn.com/problems/unique-email-addresses

func numUniqueEmails(emails []string) int {
	m := map[string]struct{}{}
	for _, email := range emails {
		l1 := len(email)
		var i int
		var cur []byte
		for i < l1 && email[i] != '@' {
			if email[i] == '+' {
				for i < l1 && email[i] != '@' {
					i++
				}
			} else {
				if email[i] != '.' {
					cur = append(cur, email[i])
				}
				i++
			}
		}
		for i < l1 {
			cur = append(cur, email[i])
			i++
		}
		m[string(cur)] = struct{}{}
	}
	return len(m)
}
