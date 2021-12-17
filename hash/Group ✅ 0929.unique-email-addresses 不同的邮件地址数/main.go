package main

// https://leetcode-cn.com/problems/unique-email-addresses

// ❓ 不同的邮件地址数
// b.y+ig@gmail.com
// .会被忽略
// +忽略 +~@内容
func numUniqueEmails(emails []string) int {
	mailMp := map[string]struct{}{}
	for _, email := range emails {
		emailL := len(email)
		var idx int
		var emailTmp []byte
		// 计算 b.y+ig
		for idx < emailL && email[idx] != '@' {
			if email[idx] == '+' {
				// 移动到后 @gmail.com
				for idx < emailL && email[idx] != '@' {
					idx++
				}
			} else {
				if email[idx] != '.' {
					emailTmp = append(emailTmp, email[idx])
				}
				idx++
			}
		}

		// 计算 @gmail.com
		emailTmp = append(emailTmp, email[idx:]...)

		mailMp[string(emailTmp)] = struct{}{}
	}
	return len(mailMp)
}
