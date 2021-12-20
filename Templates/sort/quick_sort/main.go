package main

import (
	"fmt"
)

var arr = []int{9, 3, 6, 4, 301, -200, -300, 5, 7, 8, 1, 2, 5, 5, 4, 3, 2, 1, 3, 6, 7, 8, 132, 132, 133, 134, 135, 136, 137, 138}

//var arr = []int{2, 1, 3}

func main() {
	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("%+v", arr)
}

func quickSort(arr []int, left, right int) {
	l, r := left, right
	cur := left
	for l <= r {
		for l <= cur && arr[l] <= arr[cur] {
			l++
		}
		if l < cur {
			arr[cur], arr[l] = arr[l], arr[cur]
			cur = l
		}

		for cur <= r && arr[cur] <= arr[r] {
			r--
		}

		if cur < r {
			arr[cur], arr[r] = arr[r], arr[cur]
			cur = r
		}
	}

	if 1 < cur-left {
		quickSort(arr, left, cur-1)
	}

	if 1 < right-cur {
		quickSort(arr, cur+1, right)
	}
}
