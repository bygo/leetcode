package main

import (
	"sort"
	"strings"
)

// https://leetcode.cn/problems/delete-duplicate-folders-in-system

type trie struct {
	strMp map[string]*trie
	val   string
	del   bool
}

func deleteDuplicateFolder(paths [][]string) [][]string {
	root := &trie{}
	for _, path := range paths {
		node := root
		for _, str := range path {
			if node.strMp == nil {
				node.strMp = map[string]*trie{}
			}
			if node.strMp[str] == nil {
				node.strMp[str] = &trie{}
			}
			node = node.strMp[str]
			node.val = str
		}
	}

	strMpTries := map[string][]*trie{}
	var build func(node *trie) string
	build = func(node *trie) string {
		if node.strMp == nil {
			return "(" + node.val + ")"
		}
		expr := []string{}
		for _, n := range node.strMp {
			expr = append(expr, build(n))
		}
		sort.Strings(expr)
		subExpr := strings.Join(expr, "")
		strMpTries[subExpr] = append(strMpTries[subExpr], node)
		return "(" + node.val + subExpr + ")"
	}
	build(root)

	for _, tries := range strMpTries {
		if 1 < len(tries) {
			for _, f := range tries {
				f.del = true
			}
		}
	}

	var pathsRes [][]string
	path := []string{}
	var search func(*trie)
	search = func(node *trie) {
		if node.del {
			return
		}
		path = append(path, node.val)
		pathsRes = append(pathsRes, append([]string(nil), path...))
		for _, n := range node.strMp {
			search(n)
		}
		path = path[:len(path)-1]
	}
	for _, node := range root.strMp {
		search(node)
	}
	return pathsRes
}
