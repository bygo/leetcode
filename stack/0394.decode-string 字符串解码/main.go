package main

//Link: https://leetcode-cn.com/problems/decode-string

func decodeString(s string) string {
	var stack = make([][]rune, 0)
	var nums = make([]int, 0)
	var cur []rune
	var num int
	for _, v := range s {
		if 'a' <= v && v <= 'z' || 'A' <= v && v <= 'Z' {
			cur = append(cur, v)
		} else if '0' <= v && v <= '9' {
			num = num*10 + int(v-'0')
		} else if v == '[' {
			nums = append(nums, num)
			num = 0
			stack = append(stack, cur)
			cur = []rune{}
		} else {
			top := len(stack) - 1
			n := nums[top]
			nums = nums[:top]
			elem := cur
			for 1 < n {
				cur = append(cur, elem...)
				n--
			}
			cur = append(stack[top], cur...)
			stack = stack[:top]
		}
	}
	return string(cur)
}
