package main

//Title: Intersection of Three Sorted Arrays
// Link: https://leetcode-cn.com/problems/intersection-of-three-sorted-arrays

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	l1, l2, l3 := len(arr1), len(arr2), len(arr3)
	var i, j, k int
	var res []int
	for i < l1 && j < l2 && k < l3 {
		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			res = append(res, arr1[i])
			i++
			j++
			k++
		} else if arr1[i] < arr2[k] {
			i++
		} else if arr2[j] < arr3[k] {
			j++
		} else {
			k++
		}
	}
	return res
}
