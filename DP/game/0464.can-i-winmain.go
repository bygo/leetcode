package game

// https://leetcode.cn/problems/can-i-win

func canIWin(maxChoosableInteger, desiredTotal int) bool {
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	f := make([]int8, 1<<maxChoosableInteger)
	for i := range f {
		f[i] = -1
	}
	var dfs func(int, int) int8
	dfs = func(used, curTot int) int8 {
		if f[used] != -1 {
			return f[used]
		}
		for idx := 0; idx < maxChoosableInteger; idx++ {
			num := idx + 1
			if used>>idx&1 == 1 {
				continue
			}

			if curTot+num < desiredTotal && dfs(used|1<<idx, curTot+num) == 1 {
				continue
			}
			f[used] = 1
			return f[used]
		}
		f[used] = 0
		return f[used]
	}
	return dfs(0, 0) == 1
}
