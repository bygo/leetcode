package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/design-in-memory-file-system

func main() {
	fs := FileSystem{root: &file{files: map[string]*file{}}}
	fs.Mkdir("/a/b/c")
	fs.AddContentToFile("/a/b/c/d", "hello")
	fmt.Printf("%+v", fs.Ls("/"))
}

type file struct {
	isFile  bool
	content []byte
	files   map[string]*file
}
type FileSystem struct {
	root *file
}

func Constructor() FileSystem {
	return FileSystem{
		root: &file{files: map[string]*file{}},
	}
}

func (fs *FileSystem) Ls(path string) []string {
	dirs := Split(path)
	root := fs.root
	for _, dir := range dirs {
		println(dir)
		if root.files[dir] == nil {
			return nil
		}
		root = root.files[dir]
	}

	if root.isFile {
		return []string{dirs[len(dirs)-1]}
	}
	var files []string
	for f := range root.files {
		files = append(files, f)
	}
	sort.Strings(files)
	return files
}

func (fs *FileSystem) Mkdir(path string) *file {
	dirs := Split(path)
	root := fs.root
	for _, dir := range dirs {
		if root.files[dir] == nil {
			root.files[dir] = &file{files: map[string]*file{}}
		}
		root = root.files[dir]
	}
	return root
}

func (fs *FileSystem) AddContentToFile(filePath string, content string) {
	root := fs.Mkdir(filePath)
	root.isFile = true
	root.content = append(root.content, content...)
}

func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	dirs := Split(filePath)
	root := fs.root
	for _, dir := range dirs {
		root = root.files[dir]
	}
	return string(root.content)
}

func Split(path string) []string {
	var dirs []string
	var bufCur []byte
	pathL := len(path)
	for i := 1; i < pathL; i++ {
		ch := path[i]
		if ch == '/' {
			dirs = append(dirs, string(bufCur))
			bufCur = []byte{}
		} else {
			bufCur = append(bufCur, ch)
		}
	}
	if 0 < len(bufCur) {
		dirs = append(dirs, string(bufCur))
	}
	return dirs
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */
