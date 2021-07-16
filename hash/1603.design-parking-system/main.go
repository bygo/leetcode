package main

//Title: Design Parking System
// https://leetcode-cn.com/problems/design-parking-system

type ParkingSystem struct {
	Packing int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Packing: big<<10 + medium<<20 + small<<30,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.Packing&(1023<<(carType*10)) == 0 {
		return false
	}
	this.Packing = this.Packing - 1<<(carType*10)
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
