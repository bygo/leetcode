package main

// https://leetcode.cn/problems/stone-game

func stoneGame(piles []int) bool {
	pL := len(piles)
	f := make([][]int, pL)
	for i := range f {
		f[i] = make([]int, pL)
		f[i][i] = piles[i]
	}

	for i := 0; i < pL; i++ {
		for j := 0; j < i; j++ {
			f[j][i] = max(piles[i]-f[j][i-1], piles[j]-f[j+1][i])
		}
	}

	return 0 < f[0][pL-1]
}

// 压缩
func stoneGame(piles []int) bool {
	pL := len(piles)
	top := pL - 1
	f := make([]int, pL)
	f[0] = piles[0]

	for r := 0; r < top; r++ {
		for l := 0; l < r; l++ {
			f[l] = max(piles[l]-f[l+1], piles[r]-f[r-1])
		}
	}
	return 0 < max(piles[0]-f[1], piles[top]-f[top-1])
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
