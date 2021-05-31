package main

//Link: https://leetcode-cn.com/problems/simplify-path

// 状态机
func simplifyPath(path string) string {
	if path == "" {
		return "."
	}
	var stack []string
	var cur []rune
	path += "/"
	for _, v := range path {
		if v == '/' {
			l := len(cur)
			if 2 == l && cur[0] == '.' && cur[1] == '.' {
				if 0 < len(stack) {
					stack = stack[:len(stack)-1]
				}
			} else if 1 == l && cur[0] != '.' || 1 < l {
				stack = append(stack, string(cur))
			}
			cur = []rune{}
		} else {
			cur = append(cur, v)
		}
	}
	if 0 == len(stack) {
		return "/"
	}
	var res []rune
	for _, s := range stack {
		res = append(res, '/')
		for _, v := range s {
			res = append(res, v)
		}
	}

	return string(res)
}
