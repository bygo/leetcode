package main

import (
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/number-of-atoms

// "K4(ON(SO3)2)2"

func countOfAtoms(formula string) string {
	var stack = []map[string]int{{}}
	var i = 0
	var n = len(formula)
	parseNum := func() int {
		if i == n || formula[i] < '0' || '9' < formula[i] {
			return 1
		}
		var num int
		for i < n && '0' <= formula[i] && formula[i] <= '9' {
			num = num*10 + int(formula[i]-'0')
			i++
		}
		return num
	}
	for i < n {
		switch formula[i] {
		case '(':
			stack = append(stack, map[string]int{})
			i++
		case ')':
			top := len(stack) - 1
			i++
			num := parseNum()
			for k, v := range stack[top] { // 如果有数字 整体倍增  并入上一栈
				stack[top-1][k] += v * num
			}
			stack = stack[:top]
		default:
			left := i
			i++
			for i < n && 'a' <= formula[i] && formula[i] <= 'z' {
				i++
			}
			stack[len(stack)-1][formula[left:i]] += parseNum()
		}
	}

	var keys []string
	for k := range stack[0] {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var res []byte
	for _, k := range keys {
		res = append(res, []byte(k)...)
		num := stack[0][k]
		if 1 < num {
			res = append(res, strconv.Itoa(stack[0][k])...)
		}
	}
	return string(res)
}
