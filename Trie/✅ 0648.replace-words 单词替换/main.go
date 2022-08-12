package main

// https://leetcode.cn/problems/replace-words

// ❓ 单词替换词根

func replaceWords(dictionary []string, sentence string) string {
	root := trie{}
	for _, dict := range dictionary {
		root.insert(dict)
	}
	var bufSentence []byte
	var bufCur []byte
	var fn = func() {
		strSearch := root.search(string(bufCur))
		if 0 < len(strSearch) {
			bufSentence = append(bufSentence, strSearch...)
		} else {
			bufSentence = append(bufSentence, string(bufCur)...)
		}
	}

	for i := range sentence {
		ch := sentence[i]
		if ch == ' ' {
			fn()
			bufCur = []byte{}
			bufSentence = append(bufSentence, ' ')
		} else {
			bufCur = append(bufCur, ch)
		}
	}
	if 0 < len(bufCur) {
		fn()
	}
	return string(bufSentence)
}

type trie struct {
	chMp  [26]*trie
	isEnd bool
}

func (t *trie) insert(word string) {
	node := t
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &trie{}
		}
		node = node.chMp[ch]
	}
	node.isEnd = true
}

func (t *trie) search(word string) string {
	node := t
	bufCur := []byte{}
	for i := range word {
		ch := word[i]
		bufCur = append(bufCur, ch)
		ch -= 'a'
		if node.chMp[ch] == nil {
			return ""
		} else if node.chMp[ch].isEnd {
			return string(bufCur)
		} else {
			node = node.chMp[ch]
		}
	}
	return string(bufCur)
}
