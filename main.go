package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var isDir bool

func main() {
	nameFlg := flag.String("n", "", "file name")
	typeFlg := flag.String("t", "", "file type")

	flag.Parse()

	if *typeFlg == "d" {
		isDir = true
	}

	args := flag.Args()
	if len(args) == 0 {
		log.Fatalf("Usage: go run main.go [-options <value>] [path]")
	}
	path := args[0]

	if _, err := os.Stat(path); err != nil {
		log.Fatalf("Error: Invalid path: %s", path)
	}

	walkDir(path, *nameFlg)

}

func walkDir(root string, pattarn string) error {
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}

		err = outputFile(path, pattarn, info)
		return err
	})

	return err
}

func outputFile(path string, pattarn string, info fs.FileInfo) error {
	matched, err := regexp.MatchString(pattarn, info.Name())
	if err != nil {
		return err
	}
	if matched && info.IsDir() == isDir {
		fmt.Println(path)
	}
	return nil
}
