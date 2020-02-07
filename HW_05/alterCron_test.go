package main

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestGetOrCreateDir(t *testing.T) {
	right := []string{"HW_05", "Images"}

	dir, imageDir := right[0], right[1]

	absPath, err := filepath.Abs(dir)
	if err !=nil{
		log.Fatal(err)
	}

	fullPath := filepath.Join(absPath, imageDir)

	r := GetOrCreateDir(DirTo, ImageDir)

	if f, err := os.Stat(fullPath); os.IsExist(err){
		t.Fatalf("fileInfo: %v\nerrTest: %v\n", f, err)
	}


	testingSlice := strings.Split(r, "/")
	s := testingSlice[len(testingSlice)-2:]
	if !reflect.DeepEqual(right, s){
		t.Fatalf("Ожидается: %v\nПолучено: %v", right, s)
	}
}

func TestDelFile(t *testing.T) {

}