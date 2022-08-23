package main

// https://leetcode.cn/problems/integer-replacement

func integerReplacement(n int) int {
	var cnt int
	for n != 1 {
		switch {
		case n == 3: // 3+=1  4>>1  2 ❌
			return cnt + 2
		case n%2 == 0:
			n >>= 1
			cnt++
		case n%4 == 1: // 5 9 13 17
			//n -= 1
			n >>= 1
			cnt += 2
		case n%4 == 3:
			// 7 11 15 19 23 27 31
			// ✅ Any multiple of 4 must be able to fold in half more than once
			// ❌ It has to be odd if it's a multiple of 2
			// except 3, `n+1 & >>1` is better than `n-1`
			n += 1
			n >>= 1
			cnt += 2
		}
	}
	return cnt

}

// 32 16  4 2 1
// 30 15 14 7 6 3 2 1

// 28 14  7 6 3 2 1
// 26 13 12 6 3 2 1

// 24 12  6 3 2 1
// 22 11 10 5 4 2 1

// 20 10 5 4 2 1
// 18  9 8 4 2 1

// 16 8 4 2 1
// 14 7 6 3 2 1

// 12 6 3 2 1
// 10 5 4 2 1

// 8 4 2 1
// 6 3 2 1
