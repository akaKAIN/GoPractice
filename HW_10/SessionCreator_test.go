package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestNewFileNameInPathTo(t *testing.T) {
	path, err := filepath.Abs(".")
	CheckErr("Testing Dir-func", err)
	fl, err := os.Stat(path)
	CheckErr("", err)
	if fl.IsDir(){

	}

	fmt.Println("abs-path:", path)
	dir := filepath.Dir(path)
	fmt.Println("dir:", dir)

}
