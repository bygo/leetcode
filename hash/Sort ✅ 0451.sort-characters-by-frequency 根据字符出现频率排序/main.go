package main

import (
	"sort"
)

// https://leetcode-cn.com/problems/sort-characters-by-frequency

// ❓ 根据出现频率降序

func frequencySort(s string) string {
	// 出现的频率
	var chMpCnt = map[byte]int{}
	var cntMax = 0
	for i := range s {
		ch := s[i]
		chMpCnt[ch]++
		if cntMax < chMpCnt[ch] {
			cntMax = chMpCnt[ch]
		}
	}

	// 计数排序
	var cntsChs = make([][]byte, cntMax+1)
	for ch, cnt := range chMpCnt {
		cntsChs[cnt] = append(cntsChs[cnt], ch)
	}

	var chsSort = make([]byte, 0, len(s))
	for cnt := cntMax; 0 < cnt; cnt-- {
		for _, ch := range cntsChs[cnt] {
			for i := 0; i < cnt; i++ {
				chsSort = append(chsSort, ch)
			}
		}
	}
	return string(chsSort)
}

func frequencySortQuick(s string) string {
	var chMpCnt = map[byte]int{}
	for i := range s {
		ch := s[i]
		chMpCnt[ch]++
	}

	type pair struct {
		cnt int
		ch  byte
	}

	var pairs []pair
	for ch, cnt := range chMpCnt {
		pairs = append(pairs, pair{
			cnt: cnt,
			ch:  ch,
		})
	}
	// 直接排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[j].cnt < pairs[i].cnt
	})

	var chsSort []byte
	for _, p := range pairs {
		for i := 0; i < p.cnt; i++ {
			chsSort = append(chsSort, p.ch)
		}
	}
	return string(chsSort)
}
