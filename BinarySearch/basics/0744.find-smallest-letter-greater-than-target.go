package main

// https://leetcode.cn/problems/find-smallest-letter-greater-than-target

// general
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

// offset
func nextGreatestLetter(letters []byte, target byte) byte {
	lo, hi := -1, len(letters)-1
	for lo < hi {
		mid := int(uint(lo+hi+1) >> 1)
		if letters[mid] <= target {
			lo = mid
		} else if target < letters[mid] {
			hi = mid - 1
		}
	}
	lo += 1
	if lo == len(letters) {
		return letters[0]
	}
	return letters[lo]
}

// equal
func nextGreatestLetter(letters []byte, target byte) byte {
	lo, hi := 0, len(letters)-1
	for lo <= hi {
		mid := int(uint(lo+hi) >> 1)
		if letters[mid] <= target {
			lo = mid + 1
		} else if target < letters[mid] {
			hi = mid - 1
		}
	}
	if lo == len(letters) {
		return letters[0]
	}

	return letters[lo]
}
