package main

// https://leetcode.cn/problems/design-file-system

// ❓ 设计文件系统

type FileSystem struct {
	pathMpTrie map[string]*FileSystem
	val        int
}

func Constructor() FileSystem {
	return FileSystem{pathMpTrie: map[string]*FileSystem{}}
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

func (fs *FileSystem) CreatePath(path string, value int) bool {
	node := fs
	paths := toPaths(path)
	pathsL := len(paths)
	if pathsL == 0 {
		return false
	}
	pathsTop := pathsL - 1
	// parent
	for i := 0; i < pathsTop; i++ {
		pathStr := paths[i]
		if node.pathMpTrie[pathStr] == nil {
			return false
		}
		node = node.pathMpTrie[pathStr]
	}

	pathStrTop := paths[pathsTop]

	if node.pathMpTrie[pathStrTop] == nil {
		node.pathMpTrie[pathStrTop] = &FileSystem{
			pathMpTrie: map[string]*FileSystem{},
			val:        value,
		}
		return true
	}
	return false
}

func (fs *FileSystem) Get(path string) int {
	paths := toPaths(path)
	pathsL := len(paths)
	if pathsL == 0 {
		return -1
	}

	node := fs
	for i := 0; i < pathsL; i++ {
		pathStr := paths[i]
		if node.pathMpTrie[pathStr] == nil {
			return -1
		}
		node = node.pathMpTrie[pathStr]
	}

	return node.val
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.CreatePath(path,value);
 * param_2 := obj.Get(path);
 */
