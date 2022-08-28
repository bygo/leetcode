package main

// https://leetcode.cn/problems/bulb-switcher-ii

func flipLights(n int, presses int) int {
	if presses == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	if n == 2 {
		if presses == 1 { // 00 01 10 10
			return 3
		} else {
			return 4
		}
	}
	// 3 <= n
	if presses == 1 { // 000 010 101 100
		return 4
	} else if presses == 2 { // 010 101 100 000 110 001 111
		return 7
	} else { // 3 <= presses 011
		return 8
	}

}
