package main

import (
	"archive/zip"
	"os"
	"path/filepath"
)

var (
	global_variable_maybe *zip.Writer
)

func walkAction(path string, info os.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if filepath.Ext(path) == ".xml" {
		global_variable_maybe.Create(path)
		println(path)
	}
	return nil
}

func main() {

	f, _ := os.Create("test.zip")
	defer f.Close()

	global_variable_maybe = zip.NewWriter(f)

	err := filepath.WalkDir(".", walkAction)
	if err != nil {
		println(err)
		return
	}
}
