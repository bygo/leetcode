package main

// https://leetcode.cn/problems/remove-duplicate-letters

func removeDuplicateLetters(s string) string {
	counter := map[byte]int{}
	for i := range s { // 可以出现的次数
		counter[s[i]] += 1
	}

	is := map[byte]bool{} // 是否在栈内
	var stack []byte      // 伪 单调递增栈
	for i := range s {
		b := s[i]
		if !is[b] {
			top := len(stack) - 1
			for 0 <= top && b < stack[top] {
				if counter[stack[top]] == 0 { // 后续不再有本值
					break
				}

				is[stack[top]] = false
				stack = stack[:top]
				top--
			}
			stack = append(stack, b)
			is[b] = true
		}
		counter[b]--
	}

	return string(stack)
}
