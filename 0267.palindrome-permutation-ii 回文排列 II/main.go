package main

// https://leetcode-cn.com/problems/palindrome-permutation-ii
func generatePalindromes(s string) []string {
	m := map[byte]int{}
	l1 := len(s)
	for i := 0; i < l1; i++ {
		m[s[i]]++
	}

	var odd int

	var b []byte
	for k, num := range m {
		if num%2 == 1 {
			odd += 1
		} else {
			for i := 0; i < num/2; i++ {
				b = append(b, k)
			}
		}
	}

	l2 := len(b)
	var res []string
	if 1 < odd {
		return res
	}

	var cur []byte
	var dfs func()
	m2 := map[string]struct{}{}
	dfs = func() {
		if len(cur) == l2 {
			m2[string(cur)] = struct{}{}
		}
	}
}
