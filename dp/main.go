package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var md bytes.Buffer

	all, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range all {
		sqlDir, _ := ioutil.ReadDir(fmt.Sprintf("./%s", item.Name()))
		var b = bytes.NewBufferString(fmt.Sprintf("# %s \n", item.Name()))
		for _, sqlFile := range sqlDir {
			b.WriteString("```go\n")
			s, _ := ioutil.ReadFile(fmt.Sprintf("./%s/%s", item.Name(), sqlFile.Name()))
			b.Write(s)
			b.WriteString("\n```\n\n")
			md.WriteString(b.String())
		}
	}

	ioutil.WriteFile("readme.md", md.Bytes(), 0777)
}
