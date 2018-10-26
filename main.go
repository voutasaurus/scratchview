package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	os.Remove(os.Args[0])
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if fi.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
		return
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, file)
}
