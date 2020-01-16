package main

import (
	"sort"
	"testing"
)

var inputs = [][2][]int{
	[2][]int{[]int{1000001}, []int{1000000}},
	[2][]int{[]int{1, 3}, []int{2}},
	[2][]int{[]int{1, 3}, []int{2}},                   //2
	[2][]int{[]int{1, 2}, []int{}},                    //1.5
	[2][]int{[]int{1}, []int{2, 3}},                   //2
	[2][]int{[]int{}, []int{3}},                       //3
	[2][]int{[]int{1, 2}, []int{3, 4, 5, 6, 7, 8, 9}}, //5
	[2][]int{[]int{1, 2}, []int{-1, 3}},               //1.5
	[2][]int{[]int{1}, []int{2, 3, 4, 5, 6}},          //3.5
	[2][]int{[]int{1, 5}, []int{2, 3, 4, 6}},          //3.5
	[2][]int{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 13, 14, 15, 16}, []int{11, 12, 17, 18, 19, 20}}, //10.5
	[2][]int{[]int{1}, []int{2, 3, 4, 5, 6}},                                                      //3.5
	[2][]int{[]int{1, 2, 2}, []int{1, 2, 3}},                                                      //2
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			findMedianSortedArrays(input[0], input[1])
		}
	}
}

func Test(t *testing.T) {
	for _, input := range inputs {
		var tmp []int
		tmp = append(input[0], input[1]...)
		sort.Ints(tmp)
		l := len(tmp)
		var median float64
		if l%2 == 1 {
			median = float64(tmp[l/2])
		} else {
			median = float64(tmp[l/2]+tmp[l/2-1]) / 2
		}
		if findMedianSortedArrays(input[0], input[1]) != float64(median) {
			t.Error("Error")
		}
	}

}
