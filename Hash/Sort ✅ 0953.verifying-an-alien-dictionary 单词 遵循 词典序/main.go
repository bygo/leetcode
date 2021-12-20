package main

// https://leetcode-cn.com/problems/verifying-an-alien-dictionary

func isAlienSorted(words []string, order string) bool {
	m := [26]int{}
	for i, v := range order {
		m[int(v-'a')] = i
	}

	var l, r string
	ll := 0
	for i := range words {
		r = words[i]
		rl := len(r)
		j := 0
		for j < ll && j < rl {
			p1 := m[l[j]-'a']
			p2 := m[r[j]-'a']
			if p1 < p2 {
				break
			} else if p2 < p1 {
				return false
			} else if p1 == p2 {
				j++
			}
		}
		if j == rl && rl < ll {
			return false
		}
		l = r
		ll = rl
	}
	return true
}
