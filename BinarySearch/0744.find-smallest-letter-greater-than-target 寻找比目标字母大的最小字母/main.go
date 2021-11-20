package main

// https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)
	for l < r {
		mid := l + (r-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l == len(letters) {
		return letters[0]
	} else {
		return letters[l]
	}
}
