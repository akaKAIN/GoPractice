package main

import (
	"hash/crc32"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
	"time"
)

func CreateTestFile(fileName string, t *testing.T) *os.FileInfo {
	/*Функция создания файла для проведения тестирования*/

	//Проверяем, что файла с таким же именем не существует.
	_, err := os.Stat(fileName)
	if err == nil {
		t.Fatalf("Файл %q уже существует.", fileName)
	}

	//Создаем файл
	if err := ioutil.WriteFile(fileName, []byte("test text"), os.ModePerm); err != nil {
		t.Fatalf("Ошибка создания тестового файла %q", fileName)
	}

	//Проверяем, что файл успешно создан
	NewFile, _ := os.Stat(fileName)
	if NewFile.Name() != fileName {
		t.Fatalf("Названия файлов не совпадают.\nОжидается: %s\nСоздан: %s", fileName, NewFile.Name())
	}
	return &NewFile
}

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
	for _, file := range FileSlice{
		if err := ioutil.WriteFile(file, []byte("test"), os.ModePerm); err != nil{
			t.Fatal(err)
		}
	}
	for i := range right{
		FI, err := os.Stat(FileSlice[i])
		if err != nil {
			t.Fatal(err, FileSlice[i])
		}
		if right[i] != CheckFile(FI) {
			t.Fatalf("Файл %q: %t", FI.Name(), right[i])
		}
	}

	//Удаляем созданные для теста файлы.
	for _, file := range FileSlice{
		if err := os.Remove(file); err !=nil {
			t.Fatal(err)
		}
	}
}

//func TestCreateCopyFile(t *testing.T) {
//StartLoop:
//	fileName := strconv.Itoa(int(time.Now().UnixNano()))
//
//	//Проверяем не существует ли уже тестовый файл.
//	if _, err := os.Stat(fileName); err == nil {
//		t.Logf("Тестовый файл %q уже существует: %s", fileName, err)
//		goto StartLoop
//	}
//	absBasePath, err := filepath.Abs(".")
//	if err != nil {
//		t.Fatal(err)
//	}
//	fileNamePath := filepath.Join(absBasePath, fileName)
//
//	//Создает тестовую директорию
//	if err := GetOrCreateDir(".", "TestingCreateFile")
//	PathToTestFile := filepath.Join(TestStorage, fileName)
//
//	//Отложенна функция очистки тестовой директории от содержимого
//	defer func() {
//		if _, err := os.Stat(TestStorage); err != nil {
//			t.Fatal(err)
//		}
//		if err := os.RemoveAll(TestStorage); err != nil {
//			t.Fatal(err)
//		}
//	}()
//
//	//Проверяем, что в ней нет тестируемого файла
//	if _, err := os.Stat(PathToTestFile); err == nil {
//		t.Fatal(err)
//	}
//
//	//Создаем тестовый файл, затем делаем отложенное удаление файла.
//	baseFile := *CreateTestFile(fileName, t)
//	defer func() {
//		if err := DelFile(baseFile, "."); err != nil {
//			t.Fatal(err)
//		}
//	}()
//
//	// Дописываем в базовый файл произвольную строку
//	text := fmt.Sprintf("Date of writing is %v", time.Now())
//	fb, err := os.OpenFile(baseFile.Name(), os.O_RDWR, os.ModePerm)
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer func() {
//		if err := fb.Close(); err != nil {
//			t.Fatal(err)
//		}
//	}()
//
//	//Пишем в файл
//	if _, err := fb.WriteString(text); err != nil {
//		t.Fatal(err)
//	}
//
//	//Создание и получение тестируемой копии файла в директории.
//	if err := CreateCopyFile(baseFile, ".", TestStorage); err != nil {
//		t.Fatal(err)
//	}
//
//	//Сравниваем хэши исходного файла и скопированного
//	h1 := GetFileHash(fileNamePath, t)
//	h2 := GetFileHash(PathToTestFile, t)
//	if h1 != h2 {
//		t.Fatal()
//	}
//
//}

//func TestGetFileList(t *testing.T) {
//
//}
