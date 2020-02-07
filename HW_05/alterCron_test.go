package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestGetOrCreateDir(t *testing.T) {
	right := []string{"HW_05", "Images"}

	dir, imageDir := right[0], right[1]

	absPath, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}

	fullPath := filepath.Join(absPath, imageDir)

	r := GetOrCreateDir(DirTo, ImageDir)

	if f, err := os.Stat(fullPath); os.IsExist(err) {
		t.Fatalf("fileInfo: %v\nerrTest: %v\n", f, err)
	}

	testingSlice := strings.Split(r, "/")
	s := testingSlice[len(testingSlice)-2:]
	if !reflect.DeepEqual(right, s) {
		t.Fatalf("Ожидается: %v\nПолучено: %v", right, s)
	}
}

func TestDelFile(t *testing.T) {
	fileName := strconv.Itoa(int(time.Now().UnixNano()))
	var buf = []byte("Create for test, delete file if u find it")

	if err := ioutil.WriteFile(fileName, buf, os.ModePerm); err != nil {
		t.Fatal("Ошибка создания тестового файла.")
	}
	TestFile, err := os.Stat(fileName)
	if err != nil || TestFile.IsDir() {
		t.Fatal("После тестового создания, файл не обнаружен")
	}
	if err := DelFile(TestFile, "."); err != nil {
		t.Fatalf("Ошибка удаления файла: %s", TestFile.Name())
	}
	if _, err := os.Stat(TestFile.Name()); err == nil{
		t.Fatalf("Файл %q не был удален в ходе теста", TestFile.Name())
	}
}
