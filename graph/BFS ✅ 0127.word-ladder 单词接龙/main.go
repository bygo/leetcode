package main

// https://leetcode-cn.com/problems/word-ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordToId := map[string]int{}
	graph := [][]int{}

	getId := func(word string) int {
		_, ok := wordToId[word]
		if !ok {
			wordToId[word] = len(wordToId)
			graph = append(graph, []int{})
		}
		return wordToId[word]
	}

	addEdge := func(word string) int {
		id1 := getId(word)
		char := []byte(word)
		for i, b := range char {
			char[i] = '*'
			id2 := getId(string(char))
			char[i] = b
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
		}
		return id1
	}
	for i := range wordList {
		addEdge(wordList[i])
	}

	beginId := addEdge(beginWord)
	endId, ok := wordToId[endWord]
	if !ok {
		return 0
	}

	dist := make([]int, len(wordToId))
	dist[beginId] = -1 << 31

	beginQueue := []int{beginId}
	for {
		cnt1 := len(beginQueue)
		if cnt1 == 0 {
			break
		}
		for _, q := range beginQueue {
			if q == endId {
				return (dist[q]+1<<31)/2 + 1
			}

			for _, next := range graph[q] {
				if dist[next] == 0 {
					dist[next] = dist[q] + 1
					beginQueue = append(beginQueue, next)
				}
			}
		}
		beginQueue = beginQueue[cnt1:]
	}
	return 0
}

func ladderLengthDouble(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}

	getId := func(word string) int {
		_, ok := wordId[word]
		if !ok {
			wordId[word] = len(wordId)
			graph = append(graph, []int{})
		}
		return wordId[word]
	}

	addEdge := func(word string) int {
		id1 := getId(word)
		char := []byte(word)
		for i, b := range char {
			char[i] = '*'
			id2 := getId(string(char))
			char[i] = b
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
		}
		return id1
	}
	for i := range wordList {
		addEdge(wordList[i])
	}

	beginId := addEdge(beginWord)
	endId, ok := wordId[endWord]
	if !ok {
		return 0
	}

	beginLevel := make([]int, len(wordId))
	beginLevel[beginId] = -1 << 31
	endLevel := make([]int, len(wordId))
	endLevel[endId] = -1 << 31

	beginQueue := []int{beginId}
	endQueue := []int{endId}
	for {
		cnt1 := len(beginQueue)
		cnt2 := len(endQueue)
		if cnt1 == 0 || cnt2 == 0 {
			break
		}
		for _, q := range beginQueue {
			if endLevel[q] < 0 {
				return (beginLevel[q]+endLevel[q]+2<<31)/2 + 1
			}

			for _, next := range graph[q] {
				if beginLevel[next] == 0 {
					beginLevel[next] = beginLevel[q] + 1
					beginQueue = append(beginQueue, next)
				}
			}
		}

		beginQueue = beginQueue[cnt1:]

		for _, q := range endQueue {
			if beginLevel[q] < 0 {
				return (beginLevel[q]+endLevel[q]+2<<31)/2 + 1
			}

			for _, next := range graph[q] {
				if endLevel[next] == 0 {
					endLevel[next] = endLevel[q] + 1
					endQueue = append(endQueue, next)
				}
			}
		}

		endQueue = endQueue[cnt2:]
	}
	return 0
}
