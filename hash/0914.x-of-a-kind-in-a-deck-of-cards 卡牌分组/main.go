package main

// https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards

func hasGroupsSizeX(deck []int) bool {
	if len(deck) <= 1 {
		return false
	}
	m := map[int]int{}
	for i := range deck {
		m[deck[i]]++
	}

	g := -1
	for i := range m {
		if 0 < m[i] {
			if g == -1 {
				g = m[i]
			} else {
				g = gcd(g, m[i])
			}
		}
	}
	return 2 <= g
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	} else {
		return gcd(y%x, x)
	}
}
