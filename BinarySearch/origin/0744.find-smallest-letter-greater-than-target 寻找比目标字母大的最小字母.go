package main

// https://leetcode.cn/problems/find-smallest-letter-greater-than-target

func nextGreatestLetter(letters []byte, target byte) byte {
	lo, hi := 0, len(letters)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if letters[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo == len(letters) { // 最大的字符 , 返回第一个
		return letters[0]
	}
	return letters[lo]
}
