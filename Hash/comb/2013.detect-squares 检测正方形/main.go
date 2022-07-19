package main

// https://leetcode.cn/problems/detect-squares

// ❓ 输入的点，能形成几个正方形

type DetectSquares struct {
	xMpYMpCnt map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{xMpYMpCnt: map[int]map[int]int{}}
}

func (ds *DetectSquares) Add(p []int) {
	x := p[0]
	y := p[1]
	if ds.xMpYMpCnt[x] == nil {
		ds.xMpYMpCnt[x] = map[int]int{}
	}
	ds.xMpYMpCnt[x][y]++
}

func (ds *DetectSquares) Count(p []int) int {
	x := p[0]
	y := p[1]
	yMpCnt := ds.xMpYMpCnt[x] // 所有x = x1 的 y 点
	var cntSquare int
	for yCur := range yMpCnt {
		dist := yCur - y // 需要的边长
		if dist != 0 {
			// 左边 右边
			yCurMpCnt := ds.xMpYMpCnt[x-dist] // x1-dis 的 y 点
			if yCurMpCnt != nil {
				cntSquare += yCurMpCnt[y] * yCurMpCnt[yCur] * yMpCnt[yCur] // 与其他三点成为正方形
			}

			yCurMpCnt = ds.xMpYMpCnt[x+dist] // x1+dis 的 y 点
			if yCurMpCnt != nil {
				cntSquare += yCurMpCnt[y] * yCurMpCnt[yCur] * yMpCnt[yCur] // 与其他三点成为正方形
			}
		}
	}
	return cntSquare
}
