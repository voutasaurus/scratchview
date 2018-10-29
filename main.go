package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	os.Remove(os.Args[0])
	cmd := ""
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	switch cmd {
	case "ls":
		ls(os.Args[2:]...)
	case "cat":
		cat(os.Args[2:]...)
	case "env":
		env()
	case "":
		ls(".")
	default:
		panic(cmd + ": no command found")
	}
}

func ls(args ...string) {
	path := "."
	if len(args) > 0 {
		path = args[0]
	}
	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if !fi.IsDir() {
		panic(path + " is not a directory")
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func cat(args ...string) {
	path := ""
	if len(args) > 0 {
		path = args[0]
	}
	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if fi.IsDir() {
		panic(path + " is a directory")
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, file)
}

func env() {
	for _, kv := range os.Environ() {
		fmt.Println(kv)
	}
}
