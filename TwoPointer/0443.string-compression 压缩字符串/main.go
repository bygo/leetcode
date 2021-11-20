package main

// https://leetcode-cn.com/problems/string-compression

func compress(chars []byte) int {
	var cnt int
	var cur = 1
	var top = len(chars) - 1
	var push = func(i int) {
		chars[cnt] = chars[i]
		cnt++
		if 1 < cur {
			var l = cnt
			for 0 < cur {
				chars[cnt] = byte(cur%10 + 48)
				cur /= 10
				cnt++
			}
			var r = cnt - 1
			for l < r {
				chars[l], chars[r] = chars[r], chars[l]
				l++
				r--
			}
		}
		cur = 1
	}

	for i := 0; i < top; i++ {
		if chars[i] == chars[i+1] {
			cur++
		} else {
			push(i)
		}
	}

	push(top)
	return cnt
}
