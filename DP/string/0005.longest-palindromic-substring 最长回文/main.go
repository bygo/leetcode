package main

// https://leetcode-cn.com/problems/longest-palindromic-substring

// f[l][r] = f[l+1][r-1]

// ❓ 最长回文
func longestPalindrome(s string) string {
	sL := len(s)
	f := make([][]bool, sL)
	for i := range f {
		f[i] = make([]bool, sL)
	}

	var lo, hi int
	for r := 0; r < sL; r++ {
		for l := 0; l < r; l++ {
			f[l][r] = s[l] == s[r] && (r-l <= 2 || f[l+1][r-1])
			if f[l][r] && hi-lo < r-l {
				hi = r
				lo = l
			}
		}
	}
	return s[lo : hi+1]
}
