package main

// https://leetcode.cn/problems/binary-gap

func binaryGap(n int) int {
	var idx int
	var last = -1
	var dist int
	for n != 0 {
		if n&1 == 1 {
			if last != -1 {
				distCur := idx - last
				if dist < distCur {
					dist = distCur
				}
			}
			last = idx
		}
		n >>= 1
		idx++
	}
	return dist
}
