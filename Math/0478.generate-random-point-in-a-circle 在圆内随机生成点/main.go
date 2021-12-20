package main

import (
	"math"
	"math/rand"
)

// https://leetcode-cn.com/problems/generate-random-point-in-a-circle

type Solution struct {
	x, y, radius float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		x:      x_center,
		y:      y_center,
		radius: radius,
	}
}

func (this *Solution) RandPoint() []float64 {
	//x := this.x - this.radius
	//y := this.y - this.radius
	//for {
	//	xp := x + rand.Float64()*this.radius*2
	//	yp := y + rand.Float64()*this.radius*2
	//	if math.Pow(xp-this.x, 2)+math.Pow(yp-this.y, 2) <= math.Pow(this.radius, 2) {
	//		return []float64{xp, yp}
	//	}
	//}
	f := rand.Float64() * 2 * math.Pi
	l := math.Sqrt(rand.Float64()) * this.radius
	x := this.x + l*math.Cos(f)
	y := this.y + l*math.Sin(f)
	return []float64{x, y}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
