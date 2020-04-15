package main

func numJewelsInStones(J string, S string) int {
	n := 0
	h := make(map[rune]bool)
	for _, v := range J {
		h[v] = true
	}
	for _, v := range S {
		if _, ok := h[v]; ok {
			n++
		}
	}
	return n
}
