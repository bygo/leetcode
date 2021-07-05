package main

// Link: https://leetcode-cn.com/problems/make-the-string-great

func makeGood(s string) string {
	var res []byte
	for i := range s {
		top := len(res) - 1
		if -1 < top && (s[i]-32 == res[top] || s[i]+32 == res[top]) {
			res = res[:top]
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
