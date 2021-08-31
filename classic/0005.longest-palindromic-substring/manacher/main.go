package main

// https://leetcode-cn.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	m := make([]byte, len(s)*2+3)
	m[0] = '^'
	for k, v := range s {
		m[k*2+1] = '#'
		m[k*2+2] = byte(v)
	}
	m[len(s)*2+1] = '#'
	m[len(s)*2+2] = '$'
	p := make([]int, len(m))
	var maxCenter, maxRight, center, right int
	for k := 1; k < len(m)-1 && maxRight < len(m)-k; k++ {
		if k < right {
			if p[center*2-k]+k < right { //镜像跳过
				continue
			}
			p[k] = right - k
		}

		for { //扩散
			p[k]++
			if m[k-p[k]] != m[k+p[k]] {
				break
			}
			center = k
			right = k + p[k]
		}

		if maxRight < p[k] {
			maxCenter = k
			maxRight = p[k]
		}
	}
	return s[(maxCenter-maxRight)/2 : (maxCenter+maxRight)/2-1]
}
