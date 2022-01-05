package main

// https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray

const (
	base = 131
)

func genPre(nums []int) ([]uint, []uint) {
	numsL := len(nums)
	var preHash = make([]uint, numsL+1)
	var preMul = make([]uint, numsL+1)
	preMul[0] = 1
	for i := 1; i <= numsL; i++ {
		bitLo := uint(nums[i-1])
		preHash[i] = preHash[i-1]*base + bitLo
		preMul[i] = preMul[i-1] * base
	}
	return preHash, preMul

}
func findLength(nums1 []int, nums2 []int) int {
	nums1L := len(nums1)
	nums2L := len(nums2)
	// 生成 前缀树

	var preHash1, preMul1 = genPre(nums1)
	var preHash2, preMul2 = genPre(nums2)

	check := func(curL int) bool {
		// nums1 hash值计算
		hashCur := preHash1[curL]
		mul := preMul1[curL-1]

		hashMpBool := map[uint]bool{hashCur: true}

		for idx := 0; idx < nums1L-curL; idx++ {
			bitHi := uint(nums1[idx])
			bitLo := uint(nums1[idx+curL])
			hashCur = hashCur - bitHi*mul // 减最高
			hashCur = hashCur * base        // 提升
			hashCur = hashCur + bitLo      // 加最低
			hashMpBool[hashCur] = true
		}

		// nums2 hash值计算
		hashCur = preHash2[curL]
		mul = preMul2[curL-1]
		// 合法
		if hashMpBool[hashCur] {
			return true
		}
		for i := 0; i < nums2L-curL; i++ {
			bitHi := uint(nums2[i]) * mul
			bitLo := uint(nums2[i+curL])
			hashCur = hashCur - bitHi
			hashCur = hashCur*base + bitLo
			// 合法
			if hashMpBool[hashCur] {
				return true
			}
		}
		return false
	}

	// 搜索 合法长度
	left, right := 1, min(nums1L, nums2L)+1
	for left < right {
		mid := int(uint(left+right) >> 1)
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
