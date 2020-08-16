package main

/**
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false

*/
var memory [][]bool

func isMatch(s string, p string) bool {
	memory = make([][]bool, len(s)+1)
	for i := 0; i < len(memory); i++ {
		memory[i] = make([]bool, len(p)+1)
	}
	return dp(0, 0, s, p)
}

func dp(i int, j int, s string, p string) bool {
	if memory[i][j] {
		return memory[i][j]
	}
	if len(p) == 0 {
		return len(s) == 0
	}

	ok := len(s) != 0 && (s[0] == p[0] || p[0] == '.') //如果相等，或者为. 匹配成功
	var res bool
	if len(p) >= 2 && p[1] == '*' { //长度大于2，获取下一个元素且为*
		//1.去除x*（p[2:]）,继续递归匹配
		//2.去除字符(s[1:]),继续递归匹配
		res = dp(i, j+2, s, p[2:]) || (ok && dp(i+1, j, s[1:], p)) //只要一种情况成功，计算匹配成功
	} else {
		//普通匹配
		res = ok && dp(i+1, j+1, s[1:], p[1:])
	}
	memory[i][j] = res
	return res
}

/**
思路：
1.遇到星号，递归匹配所有情况
2.缓存匹配到的串


todo:leetcode的测试用例没有缓存更快。可能大数组操作和变量赋值消耗更大，以后正则写多了,再来写这个的最佳实现
*/
