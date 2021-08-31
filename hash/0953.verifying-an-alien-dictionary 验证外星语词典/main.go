package main

// https://leetcode-cn.com/problems/verifying-an-alien-dictionary

func isAlienSorted(words []string, order string) bool {
	m := map[byte]int{}
	for i, v := range order {
		m[byte(v)] = i
	}

	left := ""
	leftLen := 0
	for i := range words {
		rightLen := len(words[i])
		j := 0
		for ; j < leftLen && j < rightLen; j++ {
			if m[left[j]] < m[words[i][j]] {
				break
			} else if m[words[i][j]] < m[left[j]] {
				return false
			}
		}
		if j == rightLen && rightLen < leftLen {
			return false
		}
		left = words[i]
		leftLen = rightLen
	}
	return true
}
