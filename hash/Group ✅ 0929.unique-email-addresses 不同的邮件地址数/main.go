package main

// https://leetcode-cn.com/problems/unique-email-addresses

// ❓ 子域名访问次数
// b.y+ig@gmail.com
// .会被忽略
// +忽略 +~@内容
func numUniqueEmails(emails []string) int {
	mailMp := map[string]struct{}{}
	for _, email := range emails {
		l1 := len(email)
		var idx int
		var cur []byte
		// 计算 b.y+ig
		for idx < l1 && email[idx] != '@' {
			if email[idx] == '+' {
				// 移动到后 @gmail.com
				for idx < l1 && email[idx] != '@' {
					idx++
				}
			} else {
				if email[idx] != '.' {
					cur = append(cur, email[idx])
				}
				idx++
			}
		}

		// 计算 @gmail.com
		cur = append(cur, email[idx:]...)
		mailMp[string(cur)] = struct{}{}
	}
	return len(mailMp)
}
