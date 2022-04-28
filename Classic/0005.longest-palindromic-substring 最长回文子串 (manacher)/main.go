package main

// https://leetcode-cn.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	sL := len(s)
	bufL := sL*2 + 3
	buf := make([]byte, bufL)
	buf[0] = '^'
	for k, v := range s {
		buf[k*2+1] = '#'
		buf[k*2+2] = byte(v)
	}
	buf[sL*2+1] = '#'
	buf[sL*2+2] = '$'
	p := make([]int, bufL)
	var maxCenter, maxRight, center, right int
	for idx := 1; idx < bufL-1 && maxRight < bufL-idx; idx++ {
		if idx < right {
			if p[center*2-idx]+idx < right { //镜像跳过
				continue
			}
			p[idx] = right - idx
		}

		for { //扩散
			p[idx]++
			if buf[idx-p[idx]] != buf[idx+p[idx]] {
				break
			}
			center = idx
			right = idx + p[idx]
		}

		if maxRight < p[idx] {
			maxCenter = idx
			maxRight = p[idx]
		}
	}
	return s[(maxCenter-maxRight)/2 : (maxCenter+maxRight)/2-1]
}

// 5 4 3
// 4 3 2
// 3 2 1
