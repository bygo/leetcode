package main

import (
	"math/rand"
	"time"
)

// https://leetcode-cn.com/problems/repeated-string-match

func strStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}

	var k1 int = 1000000000 + 7
	var k2 int = 1337
	rand.Seed(time.Now().Unix())
	var kMod1 int64 = int64(rand.Intn(k1)) + int64(k1)
	var kMod2 int64 = int64(rand.Intn(k2)) + int64(k2)

	var hash_needle int64 = 0
	for i := 0; i < m; i++ {
		hash_needle = (hash_needle*kMod2 + int64(needle[i])) % kMod1
	}
	var hash_haystack int64 = 0
	var extra int64 = 1
	for i := 0; i < m-1; i++ {
		hash_haystack = (hash_haystack*kMod2 + int64(haystack[i%n])) % kMod1
		extra = (extra * kMod2) % kMod1
	}
	for i := m - 1; (i - m + 1) < n; i++ {
		hash_haystack = (hash_haystack*kMod2 + int64(haystack[i%n])) % kMod1
		if hash_haystack == hash_needle {
			return i - m + 1
		}
		hash_haystack = (hash_haystack - extra*int64(haystack[(i-m+1)%n])) % kMod1
		hash_haystack = (hash_haystack + kMod1) % kMod1
	}
	return -1
}

func repeatedStringMatch(a string, b string) int {
	an, bn := len(a), len(b)
	index := strStr(a, b)
	if index == -1 {
		return -1
	}
	if an-index >= bn {
		return 1
	}
	return (bn+index-an-1)/an + 2
}
