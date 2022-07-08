package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var isDir bool

func main() {
	path := flag.String("p", "", "file name")
	name := flag.String("n", "", "file name")
	typeFlg := flag.String("t", "", "file type")

	flag.Parse()

	if *typeFlg == "d" {
		isDir = true
	}

	if _, err := os.Stat(*path); err != nil {
		log.Fatalf("Error: Invalid path: %s", *path)
	}

	walkDir(*path, *name)

}

func walkDir(root, filename string) error {
	wg := &sync.WaitGroup{}
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}

		go func() {
			wg.Add(1)
			outputFile(path, filename, info)
			wg.Done()
		}()
		return nil
	})

	wg.Wait()

	return err
}

func outputFile(path, filename string, info fs.FileInfo) {
	if strings.Index(info.Name(), filename) != -1 && info.IsDir() == isDir {
		fmt.Println(path)
	}
}
