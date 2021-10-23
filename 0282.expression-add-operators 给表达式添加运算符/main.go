package main

// https://leetcode-cn.com/problems/expression-add-operators

func addOperators(num string, target int) (res []string) {
	n := len(num)
	var back func(expr []byte, l, cur, mul int)
	back = func(expr []byte, l, cur, mul int) {
		if l == n {
			if cur == target {
				res = append(res, string(expr))
			}
			return
		}
		index := len(expr)
		if 0 < l {
			expr = append(expr, 0) // 占位，下面填充符号
		}
		// 枚举截取的数字长度（取多少位），注意数字可以是单个 0 但不能有前导零
		sum := 0
		for r := l; r < n; r++ {
			if r != l && num[l] == '0' {
				break
			}
			sum = sum*10 + int(num[r]-'0')
			expr = append(expr, num[r])
			if l == 0 { // 表达式开头不能添加符号
				back(expr, r+1, sum, sum)
			} else { // 枚举符号
				expr[index] = '+'
				back(expr, r+1, cur+sum, sum)
				expr[index] = '-'
				back(expr, r+1, cur-sum, -sum)
				expr[index] = '*'
				back(expr, r+1, cur-mul+mul*sum, mul*sum)
			}
		}
	}
	back(make([]byte, 0, n*2-1), 0, 0, 0)
	return
}
