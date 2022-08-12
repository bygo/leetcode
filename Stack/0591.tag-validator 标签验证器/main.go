package main

// https://leetcode.cn/problems/tag-validator

func isValid(code string) bool {
	if code[0] != '<' || code[1] == '!' {
		return false
	}
	i := 0
	n := len(code)
	var stack []string
	for i < n {
		if code[i] == '<' { // 进入匹配模式
			if code[i+1] == '!' { // 特殊匹配
				if n <= i+9 || code[i+2:i+9] != "[CDATA[" { //  先 [CDATA[
					return false
				}
				i += 9

				for i+2 < n && code[i:i+3] != "]]>" { //再 ]]>
					i += 1
				}
				i += 2
				if n < i {
					return false
				}
			} else { // 正常匹配
				if i+1 < n && code[i+1] == '/' {
					i += 2
				} else {
					i += 1
				}

				var tag []byte

				for i < n && 'A' <= code[i] && code[i] <= 'Z' {
					tag = append(tag, code[i])
					i += 1
				}

				if i == n || len(tag) <= 0 || 9 < len(tag) || code[i] != '>' {
					return false
				}

				if code[i-len(tag)-1] == '/' {
					top := len(stack) - 1
					if top == -1 { // 无前置匹配
						return false
					}
					if top == 0 && i+1 != n { // 首尾必须一对
						return false
					}
					v := stack[top]
					if v != string(tag) { // 匹配失败
						return false
					}
					stack = stack[:top]
				} else {
					stack = append(stack, string(tag))
				}
			}
		}
		i += 1
	}

	return len(stack) == 0
}
