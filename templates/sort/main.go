package main

import (
	"sort"
)

var arr = []int{9, 3, 6, 4, 301, 5, 7, 8, 1, 2, 4, 3, 2, 1, 3, 6, 7, 8, 132, 132, 133, 134, 135, 136, 137, 138}

func main() {
	//SelectSort(arr)
	//BubbleSort(arr)
	//InsertSort(arr)
	//ShellSort(arr)
	//CntSort(arr)
	//fmt.Printf("%+v", BucketSort(arr))
	//fmt.Printf("%+v", RadixSort(arr))
	println()
	//fmt.Printf("%+v", arr)
}

func SelectSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
		cur := arr[i]
		index := i
		for j := i + 1; j < l; j++ {
			if arr[j] < cur {
				cur = arr[j]
				index = j
			}
		}
		arr[index], arr[i] = arr[i], arr[index]
	}
	return arr
}

func BubbleSort(arr []int) {
	top := len(arr) - 1
	for i := 0; i < top; i++ {
		for j := top; i < j; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func InsertSort(arr []int) []int {
	l := len(arr)
	for i := 1; i < l; i++ {
		for j := i; 0 < j && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func ShellSort(arr []int) []int {
	l := len(arr)
	// gap := l / 2
	gap := l
	//for gap <= l/3 {
	//	gap = gap*3 + 1
	//}
	for 0 < gap {
		for g := 0; g < gap; g++ {
			for i := gap; i < l; i += gap {
				j := i
				for gap <= j && arr[j] < arr[j-gap] {
					arr[j], arr[j-gap] = arr[j-gap], arr[j]
					j -= gap
				}
			}
		}
		gap = gap * 10 / 13
		//gap = (gap - 1) / 3
	}
	return arr
}

func CntSort(arr []int) []int {
	a := [1000]int{}
	for i := range arr {
		a[arr[i]]++
	}
	var res []int
	for i := range a {
		for 0 < a[i] {
			res = append(res, i)
			a[i]--
		}
	}
	return res
}

func BucketSort(arr []int) []int {
	var min, max = arr[0], arr[0]
	for i := range arr {
		if arr[i] < min {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}
	const size = 5
	var bucket = [size][]int{}
	step := (max-min)/size + 1
	for i := range arr {
		index := (arr[i] - min) / step
		bucket[index] = append(bucket[index], arr[i])
	}

	var res []int
	for i := range bucket {
		sort.Ints(bucket[i])
		res = append(res, bucket[i]...)
	}

	return res
}

func RadixSort(arr []int) []int {
	bucket := [10][]int{}
	var maxDeep int
	var max = arr[0]
	for i := range arr {
		if max < arr[i] {
			max = arr[i]
		}
	}

	for 0 < max {
		max /= 10
		maxDeep++
	}

	var digit = 1
	var mod = 10
	for d := 0; d < maxDeep; d++ {
		for i := range arr {
			index := (arr[i] % mod) / digit
			bucket[index] = append(bucket[index], arr[i])
		}
		arr = []int{}
		for i := range bucket {
			for j := range bucket[i] {
				arr = append(arr, bucket[i][j])
			}
		}
		bucket = [10][]int{}

		digit *= 10
		mod *= 10
	}
	return arr
}
