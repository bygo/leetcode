package main

// https://leetcode-cn.com/problems/ransom-note

func canConstruct(ransomNote string, magazine string) bool {
	m := map[byte]int{}
	for i := range magazine {
		m[magazine[i]]++
	}

	for i := range ransomNote {
		if m[ransomNote[i]] == 0 {
			return false
		}
		m[ransomNote[i]]--
	}
	return true
}
