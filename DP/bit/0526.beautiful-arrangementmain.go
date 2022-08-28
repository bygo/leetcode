package bit

import "math/bits"

// https://leetcode.cn/problems/beautiful-arrangement

func countArrangement(n int) int {
	f := make([]int, 1<<n)
	f[0] = 1

	for subset := 0; subset < 1<<n; subset++ {
		num := bits.OnesCount(uint(subset))
		for idx := 1; idx <= n; idx++ {
			if subset>>(idx-1)&1 == 0 {
				continue
			}
			if num%idx != 0 && idx%num != 0 {
				continue
			}
			f[subset] += f[subset^1<<(idx-1)] // TODO
		}
	}
	return f[1<<n-1]
}
