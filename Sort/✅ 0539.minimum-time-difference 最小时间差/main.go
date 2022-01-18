package main

// https://leetcode-cn.com/problems/minimum-time-difference

func findMinDifference(timePoints []string) int {
	if 1440 < len(timePoints) {
		return 0
	}
	numMp := [1441]bool{}
	for _, point := range timePoints {
		var t [2]int
		var idx int
		for i := range point {
			ch := point[i]
			if ch == ':' {
				idx++
			} else {
				t[idx] = t[idx]*10 + int(point[i]-'0')
			}
		}
		num := t[0]*60 + t[1]
		if numMp[num] {
			return 0
		}
		numMp[num] = true
	}

	var numHead, numTail int
	var numRes = 1<<63 - 1
	var numPre = -1
	for num, b := range numMp {
		if !b {
			continue
		}
		if numPre == -1 {
			numHead = num
			numPre = num
		} else {
			numCur := num - numPre
			numPre = num
			if numCur < numRes {
				numRes = numCur
			}
		}
	}

	for num := 1440; 0 <= num; num-- {
		if numMp[num] {
			numTail = num
			break
		}
	}
	numCur := numHead + 1440 - numTail
	if numCur < numRes {
		numRes = numCur
	}
	return numRes
}
