package main

// https://leetcode-cn.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters

// ❓ 替换问号

func modifyString(s string) string {
	bufStr := []byte(s)
	sL := len(s)
	for i := range bufStr {
		ch := bufStr[i]
		if ch == '?' {
			var b byte = 'a'
			for ; b <= 'c'; b++ {
				if 0 < i && bufStr[i-1] == b { // 与前置相等
					continue
				} else if i < sL-1 && bufStr[i+1] == b { // 与后置相等
					continue
				} else {
					bufStr[i] = b
					break
				}
			}
		}
	}
	return string(bufStr)
}
