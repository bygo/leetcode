package main

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/unique-word-abbreviation

type ValidWordAbbr struct {
	m map[string]map[string]int
}

func Constructor(dictionary []string) ValidWordAbbr {
	m := map[string]map[string]int{}
	for i := range dictionary {
		s := Abbr(dictionary[i])
		if m[s] == nil {
			m[s] = map[string]int{}
		}
		m[s][dictionary[i]] ++
	}
	return ValidWordAbbr{
		m: m,
	}
}

func (vwa *ValidWordAbbr) IsUnique(word string) bool {
	s := Abbr(word)
	return vwa.m[s] == nil || len(vwa.m[s]) == 1 && 0 < vwa.m[s][word]
}

func Abbr(word string) string {
	l := len(word)
	b := strings.Builder{}
	b.WriteByte(word[0])
	b.WriteString(strconv.Itoa(l - 2))
	b.WriteByte(word[l-1])
	return b.String()
}
