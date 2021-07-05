package main

import (
	"bytes"
	"sort"
)

// Link: https://leetcode-cn.com/problems/sort-characters-by-frequency

func frequencySort(s string) string {
	var m = map[byte]int{}
	for i := range s {
		m[s[i]]++
	}

	type pair struct {
		count int
		char  byte
	}

	var pairs []pair
	for i, v := range m {
		pairs = append(pairs, pair{
			count: v,
			char:  i,
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[j].count < pairs[i].count
	})

	var res []byte
	for _, p := range pairs {
		res = append(res, bytes.Repeat([]byte{p.char}, p.count)...)
	}
	return string(res)
}

func frequencySort(s string) string {
	count := map[byte]int{}
	max := 0
	for i := range s {
		count[s[i]]++
		if max < count[s[i]] {
			max = count[s[i]]
		}
	}

	buckets := make([][]byte, max+1)
	for ch, c := range count {
		buckets[c] = append(buckets[c], ch)
	}

	res := make([]byte, 0, len(s))
	for i := max; 0 < i; i-- {
		for _, ch := range buckets[i] {
			res = append(res, bytes.Repeat([]byte{ch}, i)...)
		}
	}
	return string(res)
}
