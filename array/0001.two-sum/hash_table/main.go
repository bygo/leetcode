package main

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))
	for rawK, rawV := range nums {
		if hashK, ok := hashMap[target-rawV]; ok {
			return []int{rawK, hashK}
		}
		hashMap[rawV] = rawK
	}
	return nil
}
