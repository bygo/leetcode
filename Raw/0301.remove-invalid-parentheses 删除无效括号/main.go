package main

// https://leetcode-cn.com/problems/remove-invalid-parentheses

func dfs(ans *[]string, str string, start, lcount, rcount, lremove, rremove int) {
	if lremove == 0 && rremove == 0 {
		if isValid(str) {
			*ans = append(*ans, str)
		}
		return
	}

	for i := start; i < len(str); i++ {
		if i != start && str[i] == str[i-1] {
			continue
		}
		// 如果剩余的字符无法满足去掉的数量要求，直接返回
		if lremove+rremove > len(str)-i {
			return
		}
		// 尝试去掉一个左括号
		if lremove > 0 && str[i] == '(' {
			dfs(ans, str[:i]+str[i+1:], i, lcount, rcount, lremove-1, rremove)
		}
		// 尝试去掉一个右括号
		if rremove > 0 && str[i] == ')' {
			dfs(ans, str[:i]+str[i+1:], i, lcount, rcount, lremove, rremove-1)
		}
		if str[i] == ')' {
			lcount++
		} else if str[i] == ')' {
			rcount++
		}
		// 当前右括号的数量大于左括号的数量则为非法,直接返回.
		if rcount > lcount {
			break
		}
	}
}

func removeInvalidParentheses(s string) (ans []string) {
	lremove, rremove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lremove++
		} else if ch == ')' {
			if lremove == 0 {
				rremove++
			} else {
				lremove--
			}
		}
	}

	dfs(&ans, s, 0, 0, 0, lremove, rremove)
	return
}

func isValid(str string) bool {
	cnt := 0
	for _, ch := range str {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func removeInvalidParentheses(s string) []string {
	var res []string
	queue := map[string]struct{}{s: {}}
	for {
		for str := range queue {
			if isValid(str) {
				res = append(res, str)
			}
		}
		if 0 < len(res) {
			return res
		}
		cur := map[string]struct{}{}
		for str := range queue {
			for i, ch := range str {
				if 0 < i && byte(ch) == str[i-1] { // 等于上一个 所以当作已经删除了
					continue
				}
				if ch == '(' || ch == ')' {
					cur[str[:i]+str[i+1:]] = struct{}{}
				}
			}
		}
		queue = cur
	}
}
