package main

func minArray(numbers []int) int {
	l, r := 0, len(numbers)
	for l < r {
		mid := (l + r) >> 1
		cur := numbers[mid]
		if numbers[0] < cur {
			l = mid + 1
		} else if cur <= numbers[0] {
			r = mid
		}
	}

	if numbers[r] < numbers[0] {
		return r
	}
	return 0
}
