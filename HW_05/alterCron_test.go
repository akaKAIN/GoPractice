package main

import (
	"hash/crc32"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"time"
)


func GetFileHash(FilePath string, t *testing.T) uint32 {
	/*Функция возвращающая хэш указанного файла*/
	bf, err := ioutil.ReadFile(FilePath)
	if err != nil {
		t.Fatalf("HashingFunc: %v", err)
	}
	h := crc32.NewIEEE()
	if _, err := h.Write(bf); err != nil {
		t.Fatalf("HashingFuncWrite: %v", err)
	}
	return h.Sum32()
}

func TestGetOrCreateDir(t *testing.T) {

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
		t.Fatalf("Ошибка удаления файла: %q", TestFile.Name())
	}
	if _, err := os.Stat(TestFile.Name()); err == nil {
		t.Fatalf("Файл %q не был удален в ходе теста", TestFile.Name())
	}
}

func TestCheckFile(t *testing.T) {
	var FileSlice = [5]string{
		"image.exe",
		"icon.jpeg",
		"Jpeg.jpg",
		"image.doc",
		"doc.png",
	}
	var right = [5]bool{false, true, true, false, true}

	//Создаем файлы из массива
	for _, file := range FileSlice {
		if err := ioutil.WriteFile(file, []byte("test"), os.ModePerm); err != nil {
			t.Fatal(err)
		}
	}
	for i := range right {
		FI, err := os.Stat(FileSlice[i])
		if err != nil {
			t.Fatal(err, FileSlice[i])
		}
		if right[i] != CheckFile(FI) {
			t.Fatalf("Файл %q: %t", FI.Name(), right[i])
		}
	}

	//Удаляем созданные для теста файлы.
	for _, file := range FileSlice {
		if err := os.Remove(file); err != nil {
			t.Fatal(err)
		}
	}
}

func TestCreateCopyFile(t *testing.T){
	var dir = "TestDir"
	fileBase := strconv.Itoa(int(time.Now().UnixNano()))
	PathToTestFile := filepath.Join(dir, fileBase)

	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatal(err)
		}
	}()

	//Создаем файл для копирования
	if err := ioutil.WriteFile(fileBase, []byte("test text"), os.ModePerm); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Remove(fileBase); err != nil {
			t.Fatal(err)
		}
	}()

	baseFI, err := os.Stat(fileBase)
	if err != nil {
		t.Fatal(err)
	}

	if err := CreateCopyFile(baseFI, ".", dir); err != nil {
		t.Fatal(err)
	}


	//Сравниваем хэши исходного файла и скопированного
	h1 := GetFileHash(fileBase, t)
	h2 := GetFileHash(PathToTestFile, t)
	if h1 != h2 {
		t.Fatal("Файл не является копией")
	}

}

func TestGetFileList(t *testing.T) {
	var nameList = []string{"one", "two", "three", "four", "five"}
	var result = make([]string, 5)
	var dir = "TestDir"

	//Создаем тестовую папку
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	//Делаем отложенное удаление тестовой папки и всего ее содержимого
	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatal("Ошибка удаления тестовых файлов: ", err)
		}
	}()

	//Наполняем тестовую папку произвольными файлами
	for _, fileName := range nameList {
		path := filepath.Join(dir, fileName)
		if err := ioutil.WriteFile(path, []byte("test"), os.ModePerm); err != nil {
			t.Fatal(err)
		}
	}

	//Получаем отсортированный результат работы функции
	r := GetFileList(dir)

	//Получаем список имен полученных файлов
	if len(r) == 5 {
		for i, name := range r {
			result[i] = name.Name()
		}
	} else {
		t.Fatalf("\nКоличество прочитанного: %d.\nОжидалось: %d", len(r), len(nameList))
	}

	sort.Strings(nameList) //Сортируем
	if !reflect.DeepEqual(result, nameList) {
		t.Fatalf("\nПолучено: %v\nОжидалось: %v", result, nameList)
	}
}
