package main

import "strconv"

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

func (this *ValidWordAbbr) IsUnique(word string) bool {
	s := Abbr(word)
	return this.m[s] == nil || len(this.m[s]) == 1 && 0 < this.m[s][word]
}

func Abbr(word string) string {
	l := len(word)
	b := []byte{word[0]}
	b = append(b, strconv.Itoa(l-2)...)
	b = append(b, word[l-1])
	return string(b)
}

/**
 * Your ValidWordAbbr object will be instantiated and called as such:
 * obj := Constructor(dictionary);
 * param_1 := obj.IsUnique(word);
 */
