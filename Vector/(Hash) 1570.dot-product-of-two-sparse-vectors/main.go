package main

// https://leetcode.cn/problems/dot-product-of-two-sparse-vectors

// ❓ 稀疏向量的点积

type SparseVector struct {
	idxMpNum map[int]int
}

func Constructor(nums []int) SparseVector {
	idxMpNum := map[int]int{}
	for idx, num := range nums {
		// 初始化去除所有0
		if 0 < num {
			idxMpNum[idx] = num
		}
	}

	return SparseVector{
		idxMpNum: idxMpNum,
	}
}

func (sv *SparseVector) dotProduct(vec SparseVector) int {
	var sum int

	// 长度短的找长的
	if len(sv.idxMpNum) < len(vec.idxMpNum) {
		for k := range sv.idxMpNum {
			sum += sv.idxMpNum[k] * vec.idxMpNum[k]
		}
	} else {
		for k := range vec.idxMpNum {
			sum += sv.idxMpNum[k] * vec.idxMpNum[k]
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
