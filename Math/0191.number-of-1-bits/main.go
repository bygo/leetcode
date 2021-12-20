package main

func hammingWeight(num uint32) int {
	count := 0
	for 0 < num {
		if num%2 == 1 {
			count++
		}
		num /= 2
	}
	return count
}
