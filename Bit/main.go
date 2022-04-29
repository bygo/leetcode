package main

import (
	"io/ioutil"
	"os"
)

func main() {
	dirs, _ := ioutil.ReadDir("./")
	for _, dir := range dirs {
		files, _ := ioutil.ReadDir(dir.Name())
		for _, file := range files {
			if !file.IsDir() {
				continue
			}
			nameOld := file.Name()
			dirNameOld := dir.Name() + "/" + nameOld + "/main.go"
			dirNameNew := dir.Name() + "/" + nameOld + ".go"
			err := os.Rename(dirNameOld, dirNameNew)
			os.Remove(dir.Name() + "/" + nameOld)
			if err != nil {
				panic(err.Error())
			}
		}
	}
}

//func main() {
//	dirs, _ := ioutil.ReadDir("./")
//	for _, dir := range dirs {
//		files, _ := ioutil.ReadDir(dir.Name())
//		for _, file := range files {
//			nameOld := file.Name()
//			substrs := []rune{'✅', '❌', '🎆'}
//			var idx = -1
//			for _, substr := range substrs {
//				idx = strings.IndexRune(nameOld, substr)
//				if 0 <= idx {
//					break
//				}
//			}
//			if idx == -1 {
//				continue
//			}
//			dirNameOld := dir.Name() + "/" + nameOld
//			dirNameNew := dir.Name() + "/" + strings.TrimLeft(nameOld[idx+4:], " ")
//			err := os.Rename(dirNameOld, dirNameNew)
//			if err != nil {
//				panic(err.Error())
//			}
//		}
//	}
//}
