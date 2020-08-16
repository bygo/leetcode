package main

/**
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/


示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/

func letterCombinations(digits string) []string {
	m := map[rune]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	if len(digits) == 0 {
		return nil
	}
	res := make([]string, 1)
	combination := make([]string, 0)
	for _, v := range digits {
		for _, letter := range m[v] {
			for _, a := range res {
				combination = append(combination, a+string(letter))
			}
		}
		res = combination
		combination = nil
	}
	return res
}
