package main

import (
	"bytes"
	"sort"
)

// https://leetcode-cn.com/problems/sort-characters-by-frequency

func frequencySort(s string) string {
	var freq = map[byte]int{}
	var max = 0
	for i := range s {
		freq[s[i]]++
		if max < freq[s[i]] {
			max = freq[s[i]]
		}
	}

	var buckets = make([][]byte, max+1)
	for ch, c := range freq {
		buckets[c] = append(buckets[c], ch)
	}

	var res = make([]byte, 0, len(s))
	for i := max; 0 < i; i-- {
		for _, ch := range buckets[i] {
			res = append(res, bytes.Repeat([]byte{ch}, i)...)
		}
	}
	return string(res)
}

func frequencySort(s string) string {
	var freq = map[byte]int{}
	for i := range s {
		freq[s[i]]++
	}

	type pair struct {
		cnt  int
		char byte
	}

	var pairs []pair
	for i, v := range freq {
		pairs = append(pairs, pair{
			cnt:  v,
			char: i,
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[j].cnt < pairs[i].cnt
	})

	var res []byte
	for _, p := range pairs {
		res = append(res, bytes.Repeat([]byte{p.char}, p.cnt)...)
	}
	return string(res)
}
