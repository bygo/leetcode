package main

//Title: Dot Product of Two Sparse Vectors
// Link: https://leetcode-cn.com/problems/dot-product-of-two-sparse-vectors

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

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	var sum int
	if len(this.m) < len(vec.m) {
		for k, v1 := range this.m {
			if v2, ok := vec.m[k]; ok {
				sum += v1 * v2
			}
		}
	} else {
		for k, v1 := range vec.m {
			if v2, ok := this.m[k]; ok {
				sum += v1 * v2
			}
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
