package main

// https://leetcode-cn.com/problems/unique-email-addresses

// ❓ 不同的邮件地址数
// ⚠️ b.y+ig@gmail.com
// ⚠️ .会被忽略
// ⚠️ +忽略 + ~ @内容

func numUniqueEmails(emails []string) int {
	mailMp := map[string]struct{}{}
	for _, mail := range emails {
		var idx int
		var mailTmp []byte
		// 处理前半部分 b.y+ig
		for mail[idx] != '@' {
			if mail[idx] == '+' {
				// 移动到后 @gmail.com
				for mail[idx] != '@' {
					idx++
				}
			} else {
				if mail[idx] != '.' {
					mailTmp = append(mailTmp, mail[idx])
				}
				idx++
			}
		}

		// 处理后半部分 @gmail.com
		mailTmp = append(mailTmp, mail[idx:]...)

		mailMp[string(mailTmp)] = struct{}{}
	}
	return len(mailMp)
}
