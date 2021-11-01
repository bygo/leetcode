package main

// https://leetcode-cn.com/problems/dot-product-of-two-sparse-vectors

type SparseVector struct {
	m map[int]int
}

func Constructor(nums []int) SparseVector {
	m := map[int]int{}
	for k, v := range nums {
		if 0 < v {
			m[k] = v
		}
	}
	return SparseVector{
		m: m,
	}
}

func (sv *SparseVector) dotProduct(vec SparseVector) int {
	var sum int
	if len(sv.m) < len(vec.m) {
		for k := range sv.m {
			sum += sv.m[k] * vec.m[k]
		}
	} else {
		for k := range vec.m {
			sum += sv.m[k] * vec.m[k]
		}
	}
	return sum
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
