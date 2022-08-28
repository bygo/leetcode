package bit

// https://leetcode.cn/problems/partition-to-k-equal-sum-subsets

func canPartitionKSubsets(nums []int, k int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}
	if total%k != 0 {
		return false
	}
	cnt := total / k
	nL := len(nums)
	numMax := 1 << nL
	f := make([]int, numMax)
	for i := 1; i < numMax; i++ {
		f[i] = -1
	}

	for subset := 1; subset < numMax; subset++ {
		for idx, num := range nums {
			if subset>>idx&1 == 0 {
				continue
			}
			sub := subset ^ (1 << idx)
			if 0 <= f[sub] && f[sub]+num <= cnt {
				f[subset] = (f[sub] + num) % cnt
				break
			}
		}
	}
	return f[numMax-1] == 0
}
