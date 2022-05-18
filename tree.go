package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(os.Args)

	if len(os.Args) <= 1 {
		fmt.Println("tree cmd missing argument root")
		return
	}
	root := os.Args[1]
	tree(root, 0)
}

func tree(path string, level int) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	fmt.Println(getIndent(level), fi.Name())
	if !fi.IsDir() {
		return nil
	}

	fis, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, fi := range fis {
		if strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		tree(filepath.Join(path, fi.Name()), level+1)
	}
	return nil
}

func getIndent(level int) string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	return indent
}
