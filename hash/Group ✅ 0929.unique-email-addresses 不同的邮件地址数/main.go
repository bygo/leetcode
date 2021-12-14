package main

// https://leetcode-cn.com/problems/unique-email-addresses

// leet.code+ignore@gmail.com
// .会被忽略
// +忽略 +~@内容
func numUniqueEmails(emails []string) int {
	mailM := map[string]struct{}{}
	for _, email := range emails {
		l1 := len(email)
		var idx int
		var cur []byte
		// 前半部分
		for idx < l1 && email[idx] != '@' {
			if email[idx] == '+' {
				for idx < l1 && email[idx] != '@' { // 移动到后半部分
					idx++
				}
			} else {
				if email[idx] != '.' {
					cur = append(cur, email[idx])
				}
				idx++
			}
		}

		// 后半部分
		cur = append(cur, email[idx:]...)
		mailM[string(cur)] = struct{}{}
	}
	return len(mailM)
}
