package main

// https://leetcode-cn.com/problems/remove-sub-folders-from-the-filesystem

type trie struct {
	pathMp map[string]*trie
	str    string
}

func (t *trie) insert(word string) {
	paths := toPaths(word)
	node := t
	for i := range paths {
		path := paths[i]
		if node.pathMp[path] == nil {
			node.pathMp[path] = &trie{pathMp: map[string]*trie{}}
		}
		node = node.pathMp[path]
	}
	node.str = word
}

func removeSubfolders(folder []string) []string {
	root := &trie{pathMp: map[string]*trie{}}
	for _, word := range folder {
		root.insert(word)
	}

	var strs []string
	var dfs func(node *trie)
	dfs = func(node *trie) {
		if 0 < len(node.str) {
			strs = append(strs, node.str)
			return
		}
		for _, n := range node.pathMp {
			dfs(n)
		}
	}
	dfs(root)
	return strs
}

func toPaths(path string) []string {
	var buf []byte
	var paths []string
	for i := range path {
		if path[i] == '/' {
			if len(buf) != 0 {
				paths = append(paths, string(buf))
				buf = []byte{}
			}
		} else {
			buf = append(buf, path[i])
		}
	}
	if len(buf) != 0 {
		paths = append(paths, string(buf))
	}
	return paths
}
