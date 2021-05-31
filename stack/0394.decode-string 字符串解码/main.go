package main

//Link: https://leetcode-cn.com/problems/decode-string

func decodeString(s string) string {
	var stack = make([][]byte, 0)
	var nums = make([]int, 0)
	var cur []byte
	var num int
	var i = 0
	var n = len(s)
	for i < n {
		switch s[i] {
		case '[':
			nums = append(nums, num)
			num = 0
			stack = append(stack, cur)
			cur = []byte{}
			i++
		case ']':
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
			i++
		default:
			if 'a' <= s[i] && s[i] <= 'z' || 'A' <= s[i] && s[i] <= 'Z' {
				cur = append(cur, s[i])
			} else if '0' <= s[i] && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
			}
			i++
		}

	}
	return string(cur)
}
