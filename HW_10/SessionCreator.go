package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func MakeSession() (*Session, error) {
	//Проверяем, что переданы не пустые аргументы
	if From == "" || To == "" {
		err := fmt.Errorf("Неверно указан путь файла: %q -> %q", From, To)
		return nil, err
	}

	//Получаем имя, размер копируемого файла
	fileName, fileSize, err := CheckFile(From)
	if err != nil {
		err = fmt.Errorf("Ошибка копирования указанного файла (он не найден): %v", From)
		return nil, err
	}

	//Создаем файл, куда будем копировать.
CreateLoop:
	file, err := os.Create(To)
	if err != nil {
		if os.IsNotExist(err) {
			PathToFile := filepath.Dir(To)
			err := os.MkdirAll(PathToFile, os.ModePerm)
			if err != nil {
				return nil, err
			}
			goto CreateLoop

		}
		err = fmt.Errorf("Ошибка создания файла для копирования: %s", err)
		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			err := fmt.Errorf("Ошибка закрытия файла: %s", To)
			CheckErr("", err)
		}
	}()

	newFile, err := os.Stat(From)
	CheckErr("Error from get FileInfo", err)

	var s = Session{
		PathFrom: From,
		PathTo:   To,
		FileName: fileName,
		FileSize: fileSize,
		Offset:   Offset,
		Limit:    Limit,
		NewFile:  &newFile,
	}
	return &s, nil
}

func CheckFile(path string) (string, int64, error) {
	FileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", 0, err
		}
	}
	if FileInfo.IsDir() {
		err = fmt.Errorf("Не указано имя файла для копирования. Err: %s", err)
		return "", 0, err
	}
	return FileInfo.Name(), FileInfo.Size(), nil
}
