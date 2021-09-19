package main

// https://leetcode-cn.com/problems/detect-squares

type DetectSquares struct {
	m map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{m: map[int]map[int]int{}}
}

func (this *DetectSquares) Add(point []int) {
	x := point[0]
	y := point[1]
	if this.m[x] == nil {
		this.m[x] = map[int]int{}
	}
	this.m[x][y] ++
}

func (this *DetectSquares) Count(point []int) int {
	x1 := point[0]
	y1 := point[1]
	yMap := this.m[x1]
	var res int
	for y2 := range yMap {
		dis := y2 - y1
		if dis != 0 {
			leftY := this.m[x1-dis]
			if leftY != nil {
				res += leftY[y1] * leftY[y2] * yMap[y2]
			}

			rightY := this.m[x1+dis]
			if rightY != nil {
				res += rightY[y1] * rightY[y2] * yMap[y2]
			}
		}
	}
	return res
}
