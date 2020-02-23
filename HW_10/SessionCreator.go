package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func MakeSession() (*Session, error){
	if From == "" || To == "" {
		err := fmt.Errorf("Неверно указан путь файла: %q -> %q\n", From, To)
		CheckErr("", err)
		log.Fatal(err)
	}
	
	var s = new(Session)

	fileName := filepath.Base(From)
	fileOld, err := os.Open(From)
	CheckErr("OpenFile from session", err)

	if NewFileNameInPathTo {

	}
	fileNew, err := os.Create()


	return s, nil
}

func NewFileNameInPathTo ()bool{
	//Проверяем, что путь содержит в себе имя файла, куда будем копировать.
	filepath.Dir(To)
	return false
}

func CheckFile(path string) (int64, error) {
	FileInfo, err := os.Stat(path)
	if err != nil {
		switch {
		case os.IsNotExist(err):
			return 0, err
		case FileInfo.IsDir():
			err = fmt.Errorf("Не указано имя файла для копирования. Err: %s\n", err)
			return 0, err
		}
	}
	return FileInfo.Size(), nil
}