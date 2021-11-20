package main

// https://leetcode-cn.com/problems/sentence-similarity

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	m := map[string]map[string]bool{}
	for _, pair := range similarPairs {
		if m[pair[0]] == nil {
			m[pair[0]] = map[string]bool{}
		}
		if m[pair[1]] == nil {
			m[pair[1]] = map[string]bool{}
		}
		m[pair[0]][pair[1]] = true
		m[pair[1]][pair[0]] = true
	}

	if len(sentence1) != len(sentence2) {
		return false
	}
	for i := range sentence1 {
		if sentence1[i] != sentence2[i] && !m[sentence1[i]][sentence2[i]] && !m[sentence2[i]][sentence1[i]] {
			return false
		}
	}
	return true
}
