package main

// https://leetcode-cn.com/problems/find-the-difference

func findTheDifferenceMark(s string, t string) byte {
	m := [26]int{}
	for i := range s {
		ch := s[i] - 'a'
		m[ch]++
	}

	for i := range t {
		ch := t[i] - 'a'
		if m[ch] == 0 {
			return ch + 'a'
		}
		m[ch]--
	}
	return 0
}

func findTheDifferenceCount(s string, t string) byte {
	var sum = 0
	for i := range s {
		sum = sum - int(s[i]) + int(t[i])
	}
	return byte(sum + int(t[len(t)-1]))
}

func findTheDifferenceBit(s string, t string) byte {
	var sum byte
	for i := range s {
		sum ^= s[i] ^ t[i]
	}
	return sum ^ t[len(t)-1]
}
