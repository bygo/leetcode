package main

import "fmt"

// https://leetcode-cn.com/problems/word-ladder-ii

func findLadders2(beginWord string, endWord string, wordList []string) [][]string {
	var res [][]string

	l1 := len(wordList)
	m := map[string]int{}
	for i := range wordList {
		m[wordList[i]] = i
	}

	if _, ok := m[beginWord]; !ok {
		m[beginWord] = l1
		wordList = append(wordList, beginWord)
		l1++
	}

	if _, ok := m[endWord]; !ok {
		return res
	}

	check := func(a, b string) bool {
		for i := range a {
			if a[i] != b[i] {
				return a[i+1:] == b[i+1:]
			}
		}
		return false
	}

	graph := map[int]map[int]struct{}{}
	for i := range wordList {
		graph[i] = map[int]struct{}{}
	}
	for i := 0; i < l1; i++ {
		for j := i + 1; j < l1; j++ {
			if check(wordList[i], wordList[j]) {
				graph[i][j] = struct{}{}
				graph[j][i] = struct{}{}
			}
		}
	}

	dist := make([]int, len(wordList))
	dist[m[beginWord]] = -1 << 31 // 初始度

	queue := [][]int{{m[beginWord]}}
	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}

		for _, indexes := range queue {
			curIndex := indexes[len(indexes)-1]
			cur := wordList[curIndex]
			if cur == endWord {
				var tmp []string
				for _, index := range indexes {
					tmp = append(tmp, wordList[index])
				}
				res = append(res, tmp)
			} else {
				for nextIndex := range graph[curIndex] {
					if dist[nextIndex] <= dist[curIndex] { // 不可逆流，不可同级活水，所以最短路径上 进行加长的路线全部截断
						continue
					}
					dist[nextIndex] = dist[curIndex] + 1 // 入度
					next := make([]int, len(indexes))
					copy(next, indexes)
					next = append(next, nextIndex)
					queue = append(queue, next)
				}
			}
		}
		queue = queue[cnt:]
	}
	return res
}

func main() {
	//res := findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	res := findLadders("cet", "ism", []string{"kid", "tag", "pup", "ail", "tun", "woo", "erg", "luz", "brr", "gay", "sip", "kay", "per", "val", "mes", "ohs", "now", "boa", "cet", "pal", "bar", "die", "war", "hay", "eco", "pub", "lob", "rue", "fry", "lit", "rex", "jan", "cot", "bid", "ali", "pay", "col", "gum", "ger", "row", "won", "dan", "rum", "fad", "tut", "sag", "yip", "sui", "ark", "has", "zip", "fez", "own", "ump", "dis", "ads", "max", "jaw", "out", "btu", "ana", "gap", "cry", "led", "abe", "box", "ore", "pig", "fie", "toy", "fat", "cal", "lie", "noh", "sew", "ono", "tam", "flu", "mgm", "ply", "awe", "pry", "tit", "tie", "yet", "too", "tax", "jim", "san", "pan", "map", "ski", "ova", "wed", "non", "wac", "nut", "why", "bye", "lye", "oct", "old", "fin", "feb", "chi", "sap", "owl", "log", "tod", "dot", "bow", "fob", "for", "joe", "ivy", "fan", "age", "fax", "hip", "jib", "mel", "hus", "sob", "ifs", "tab", "ara", "dab", "jag", "jar", "arm", "lot", "tom", "sax", "tex", "yum", "pei", "wen", "wry", "ire", "irk", "far", "mew", "wit", "doe", "gas", "rte", "ian", "pot", "ask", "wag", "hag", "amy", "nag", "ron", "soy", "gin", "don", "tug", "fay", "vic", "boo", "nam", "ave", "buy", "sop", "but", "orb", "fen", "paw", "his", "sub", "bob", "yea", "oft", "inn", "rod", "yam", "pew", "web", "hod", "hun", "gyp", "wei", "wis", "rob", "gad", "pie", "mon", "dog", "bib", "rub", "ere", "dig", "era", "cat", "fox", "bee", "mod", "day", "apr", "vie", "nev", "jam", "pam", "new", "aye", "ani", "and", "ibm", "yap", "can", "pyx", "tar", "kin", "fog", "hum", "pip", "cup", "dye", "lyx", "jog", "nun", "par", "wan", "fey", "bus", "oak", "bad", "ats", "set", "qom", "vat", "eat", "pus", "rev", "axe", "ion", "six", "ila", "lao", "mom", "mas", "pro", "few", "opt", "poe", "art", "ash", "oar", "cap", "lop", "may", "shy", "rid", "bat", "sum", "rim", "fee", "bmw", "sky", "maj", "hue", "thy", "ava", "rap", "den", "fla", "auk", "cox", "ibo", "hey", "saw", "vim", "sec", "ltd", "you", "its", "tat", "dew", "eva", "tog", "ram", "let", "see", "zit", "maw", "nix", "ate", "gig", "rep", "owe", "ind", "hog", "eve", "sam", "zoo", "any", "dow", "cod", "bed", "vet", "ham", "sis", "hex", "via", "fir", "nod", "mao", "aug", "mum", "hoe", "bah", "hal", "keg", "hew", "zed", "tow", "gog", "ass", "dem", "who", "bet", "gos", "son", "ear", "spy", "kit", "boy", "due", "sen", "oaf", "mix", "hep", "fur", "ada", "bin", "nil", "mia", "ewe", "hit", "fix", "sad", "rib", "eye", "hop", "haw", "wax", "mid", "tad", "ken", "wad", "rye", "pap", "bog", "gut", "ito", "woe", "our", "ado", "sin", "mad", "ray", "hon", "roy", "dip", "hen", "iva", "lug", "asp", "hui", "yak", "bay", "poi", "yep", "bun", "try", "lad", "elm", "nat", "wyo", "gym", "dug", "toe", "dee", "wig", "sly", "rip", "geo", "cog", "pas", "zen", "odd", "nan", "lay", "pod", "fit", "hem", "joy", "bum", "rio", "yon", "dec", "leg", "put", "sue", "dim", "pet", "yaw", "nub", "bit", "bur", "sid", "sun", "oil", "red", "doc", "moe", "caw", "eel", "dix", "cub", "end", "gem", "off", "yew", "hug", "pop", "tub", "sgt", "lid", "pun", "ton", "sol", "din", "yup", "jab", "pea", "bug", "gag", "mil", "jig", "hub", "low", "did", "tin", "get", "gte", "sox", "lei", "mig", "fig", "lon", "use", "ban", "flo", "nov", "jut", "bag", "mir", "sty", "lap", "two", "ins", "con", "ant", "net", "tux", "ode", "stu", "mug", "cad", "nap", "gun", "fop", "tot", "sow", "sal", "sic", "ted", "wot", "del", "imp", "cob", "way", "ann", "tan", "mci", "job", "wet", "ism", "err", "him", "all", "pad", "hah", "hie", "aim", "ike", "jed", "ego", "mac", "baa", "min", "com", "ill", "was", "cab", "ago", "ina", "big", "ilk", "gal", "tap", "duh", "ola", "ran", "lab", "top", "gob", "hot", "ora", "tia", "kip", "han", "met", "hut", "she", "sac", "fed", "goo", "tee", "ell", "not", "act", "gil", "rut", "ala", "ape", "rig", "cid", "god", "duo", "lin", "aid", "gel", "awl", "lag", "elf", "liz", "ref", "aha", "fib", "oho", "tho", "her", "nor", "ace", "adz", "fun", "ned", "coo", "win", "tao", "coy", "van", "man", "pit", "guy", "foe", "hid", "mai", "sup", "jay", "hob", "mow", "jot", "are", "pol", "arc", "lax", "aft", "alb", "len", "air", "pug", "pox", "vow", "got", "meg", "zoe", "amp", "ale", "bud", "gee", "pin", "dun", "pat", "ten", "mob"})

	fmt.Printf("%+v", res)
	//"hit"
	//"cog"
	//["hot","dot","dog","lot","log","cog"]
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var res [][]string

	graph := [][]int{}

	wordToId := map[string]int{}
	idToWord := map[int]string{}
	getId := func(word string) int {
		_, ok := wordToId[word]
		if !ok {
			wordToId[word] = len(wordToId)
			idToWord[wordToId[word]] = word
			graph = append(graph, []int{})
		}
		return wordToId[word]
	}

	// 虚拟图
	addEdge := func(word string) int {
		id1 := getId(word)
		b := []byte(word)
		for i, ch := range b {
			b[i] = '*'
			id2 := getId(string(b))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			b[i] = ch
		}
		return id1
	}

	for _, s := range wordList {
		addEdge(s)
	}

	beginId, ok := wordToId[beginWord]
	if !ok {
		beginId = addEdge(beginWord)
	}
	endId, ok := wordToId[endWord]
	if !ok {
		return res
	}

	beginDist := make([]int, len(wordToId))
	beginDist[beginId] = -1 << 31
	endDist := make([]int, len(wordToId))
	endDist[endId] = -1 << 31

	beginQueue := []int{beginId}
	endQueue := []int{endId}
	bfs := func() int {
		for {
			cnt1 := len(beginQueue)
			cnt2 := len(endQueue)
			if cnt1 == 0 || cnt2 == 0 {
				break
			}
			for _, q := range beginQueue {
				if endDist[q] < 0 {
					return beginDist[q] + endDist[q] + 2<<31 + 1
				} else {
					for _, next := range graph[q] {
						if beginDist[next] == 0 {
							beginDist[next] = beginDist[q] + 1
							beginQueue = append(beginQueue, next)
						}
					}
				}
			}
			beginQueue = beginQueue[cnt1:]

			for _, q := range endQueue {
				if beginDist[q] < 0 {
					return beginDist[q] + endDist[q] + 2<<31 + 1
				} else {
					for _, next := range graph[q] {
						if endDist[next] == 0 {
							endDist[next] = endDist[q] + 1
							endQueue = append(endQueue, next)
						}
					}
				}
			}

			endQueue = endQueue[cnt2:]
		}

		return 0
	}

	length := bfs()
	if length <= 0 {
		return res
	}
	top := length - 1
	var path = []int{beginId}
	var dfs func(id int)
	dfs = func(id int) {
		if len(path) == length {
			lastId := path[top]
			cur := idToWord[lastId]
			if cur == endWord {
				tmp := []string{}
				for i := 0; i < length; i += 2 {
					tmp = append(tmp, idToWord[path[i]])
				}
				res = append(res, tmp)
			}
		} else {
			for _, i := range graph[id] {
				if beginDist[id]+1 == beginDist[i] || endDist[i]+1 == endDist[id] {
					path = append(path, i)
					dfs(i)
					path = path[:len(path)-1]
				}
			}
		}

	}

	dfs(beginId)

	return res
}
