package main

// https://leetcode.cn/problems/implement-magic-dictionary

// ❓ 实现一个魔法字典

type MagicDictionary struct {
	dictionary []string
	virMpCnt   map[string]int
	dictMpBool map[string]bool
}

func Constructor() MagicDictionary {
	return MagicDictionary{virMpCnt: map[string]int{}, dictMpBool: map[string]bool{}}
}

func (md *MagicDictionary) BuildDict(dictionary []string) {
	md.dictionary = dictionary
	for _, str := range dictionary {
		md.dictMpBool[str] = true
		bufStr := []byte(str)
		for i := range bufStr {
			ch := bufStr[i]
			bufStr[i] = '*'
			md.virMpCnt[string(bufStr)]++
			bufStr[i] = ch
		}
	}
}

func (md *MagicDictionary) Search(searchWord string) bool {
	bufSearchWord := []byte(searchWord)
	for i := range bufSearchWord {
		ch := bufSearchWord[i]
		bufSearchWord[i] = '*'
		if 2 <= md.virMpCnt[string(bufSearchWord)] || 1 == md.virMpCnt[string(bufSearchWord)] && !md.dictMpBool[searchWord] {
			return true
		}
		bufSearchWord[i] = ch
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
