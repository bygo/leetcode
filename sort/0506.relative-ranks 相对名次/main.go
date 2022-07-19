package main

import (
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/relative-ranks

type Pair struct {
	Score int
	Index int
}

var desc = []string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(score []int) []string {
	var pairs []Pair
	for i := range score {
		pairs = append(pairs, Pair{
			Score: score[i],
			Index: i,
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[j].Score < pairs[i].Score
	})

	var res = make([]string, len(pairs))

	for i := range pairs {
		if i <= 2 {
			res[pairs[i].Index] = desc[i]
		} else {
			res[pairs[i].Index] = strconv.Itoa(i + 1)
		}
	}
	return res
}
