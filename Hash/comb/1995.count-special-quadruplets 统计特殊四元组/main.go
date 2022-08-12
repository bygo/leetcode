package main

// https://leetcode.cn/problems/count-special-quadruplets

// ❓ 统计特殊四元组

//func countQuadruplets(nums []int) int {
//	numsL := len(nums)
//	var cntQuad int
//	for i := 0; i < numsL; i++ {
//		for j := i + 1; j < numsL; j++ {
//			for k := j + 1; k < numsL; k++ {
//				for l := k + 1; l < numsL; l++ {
//					if nums[i]+nums[j]+nums[k] == nums[l] {
//						cntQuad++
//					}
//				}
//			}
//		}
//	}
//	return cntQuad
//}
//
//func countQuadruplets(nums []int) int {
//	var cntQuad int
//	numsL := len(nums)
//	sumMpCnt := map[int]int{}
//	for k := numsL - 2; 2 <= k; k-- {
//		// 提前计算答案
//		l := k + 1
//		sumMpCnt[nums[l]]++
//
//		// 枚举 i j k
//		for i := 0; i < k-1; i++ {
//			for j := i + 1; j < k; j++ {
//				sum := nums[i] + nums[j] + nums[k]
//				cntQuad += sumMpCnt[sum]
//			}
//		}
//	}
//	return cntQuad
//}

func countQuadruplets(nums []int) int {
	var cntQuad int
	sumMpCnt := map[int]int{}
	numsL := len(nums)
	for j := numsL - 3; 1 <= j; j-- {
		// 提前计算 l 和 k 的差
		k := j + 1
		for l := j + 2; l < numsL; l++ {
			// j与k的组合
			sumMpCnt[nums[l]-nums[k]]++
		}

		for i := 0; i < j; i++ {
			sum := nums[i] + nums[j]
			if 0 < sumMpCnt[sum] {
				cntQuad += sumMpCnt[sum]
			}
		}
	}
	return cntQuad
}
