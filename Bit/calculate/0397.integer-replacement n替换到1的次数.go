package main

// https://leetcode-cn.com/problems/integer-replacement

// ❓ n替换到1的次数
// ⚠️ 1 <= n <= 1<<31 -1

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
			// ✅ 变成4的倍数必定可以折半多次
			// ❌ 变得2的倍数折半后必定奇数
			// 除了3 n+1 & >>1 都优于 n-1
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
