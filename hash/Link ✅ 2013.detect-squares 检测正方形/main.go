package main

// https://leetcode-cn.com/problems/detect-squares

type DetectSquares struct {
	m map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{m: map[int]map[int]int{}}
}

func (ds *DetectSquares) Add(p []int) {
	x := p[0]
	y := p[1]
	if ds.m[x] == nil {
		ds.m[x] = map[int]int{}
	}
	ds.m[x][y] ++
}

func (ds *DetectSquares) Count(p []int) int {
	x1 := p[0]
	y1 := p[1]
	yMap := ds.m[x1] // 所有x = x1 的 y 点
	var res int
	for y2 := range yMap {
		dis := y2 - y1 // 需要的边长
		if dis != 0 {
			//  左边 右边
			ys := ds.m[x1-dis] // x1-dis 的 y 点
			if ys != nil {
				res += ys[y1] * ys[y2] * yMap[y2] // 与其他三点成为正方形
			}

			ys = ds.m[x1+dis] // x1+dis 的 y 点
			if ys != nil {
				res += ys[y1] * ys[y2] * yMap[y2] // 与其他三点成为正方形
			}
		}
	}
	return res
}
