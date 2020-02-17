package main

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
 */

func isValid(s string) bool {
	m := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if mirror, ok := m[v]; ok {
			stack = append(stack, mirror)
		} else if len(stack) > 0 && stack[len(stack)-1] == v {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return true
}
