package main

// https://leetcode.cn/problems/design-parking-system

type ParkingSystem struct {
	Packing int
}

var Places = [3]int{
	1023 << 20,
	1023 << 10,
	1023,
}

var nums = [3]int{
	1 << 20,
	1 << 10,
	1,
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Packing: big<<20 + medium<<10 + small,
	}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	carType -= 1
	place := Places[carType]
	if ps.Packing&place == 0 {
		return false
	}
	ps.Packing -= nums[carType]
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
