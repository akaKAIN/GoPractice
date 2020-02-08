package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	//Инициализация меню-подсказки
	InitFlag()

	//Проверка существования директории для конечных файлов. Если ее нет - создаем.
	if err := GetOrCreateDir(DirTo); err != nil {
		log.Fatal(err)
	}
	for {
		//Получаем список файлов в базовой директории
		FileList := GetFileList(DirFrom)
		go ReplaceFiles(FileList, DirTo)

		time.Sleep(time.Duration(TIMER))
	}

}

func ReplaceFiles(FileList []os.FileInfo, PathToStorageDir string) {
	for _, file := range FileList {

		//Проверяем файл на валидность и соответствие расширения
		if CheckFile(file) {

			//Копируем файл в папку-хранилище
			err := CreateCopyFile(file, DirFrom, PathToStorageDir)
			if err != nil {
				log.Printf("CopyError: %v file:%s\n", err, file.Name())
			}

			//Удаляем старую копию файла если файл успешно скопирован
			if err == nil {
				err = DelFile(file, DirFrom)
				if err != nil {
					log.Printf("DelError: %v file:%s\n", err, file.Name())
				}
			}

			//Информирование об успешном переносе файла
			PrintInfoLog(fmt.Sprintf("Перемещен: %s", file.Name()))
		}

	}
}

func PrintInfoLog(text string) {
	if LOGGING {
		log.Println(text)
	}
}

func GetFileList(DirName string) []os.FileInfo {
	/*Возвращает список файлов указанной директории*/
	DirNamePath, err := filepath.Abs(DirName)
	if err != nil {
		log.Fatalf("Ошибка построения абсолютного пути к %q : %v", DirName, err)
	}
	files, err := ioutil.ReadDir(DirNamePath)
	if err != nil {
		log.Fatal("Ошибка чтения:", err)
	}
	return files
}

func GetOrCreateDir(dir string) error {
	/*Проверяет существует ли указанная папка для хранения. Если нет - создает такую папку.
	Возвращает полный путь до папки*/
	if _, err := os.Stat(dir); !os.IsExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			PrintInfoLog(fmt.Sprintf("Папка для хранения файлов уже существует: %v\n", dir))
			return nil
		}
		PrintInfoLog(fmt.Sprintf("Создана папка для хранения файлов по адресу: %v\n", dir))
	}
	return nil
}

func DelFile(file os.FileInfo, path string) error {
	/*Удаляет указанный файл в указанной директории*/

	PathToFile := filepath.Join(path, file.Name())
	if err := os.Remove(PathToFile); err != nil {
		return err
	}
	return nil
}

func CheckFile(file os.FileInfo) bool {
	//Сравнивает расширение переданного файла на совпадение с любым значеним из слайса с расширениями (ExtArr)
	ext := filepath.Ext(file.Name())
	for _, imgType := range ExtArr {
		if ext == imgType {
			return true
		}
	}
	return false

}

func CreateCopyFile(file os.FileInfo, dirFrom, dirTo string) error {
	/*Создает копию указанного файла в указанной директории*/

	PathFrom := filepath.Join(dirFrom, file.Name())
	PathTo := filepath.Join(dirTo, file.Name())

	buf, err := ioutil.ReadFile(PathFrom)
	if err != nil {
		PrintInfoLog(fmt.Sprintf("ОШИБКА! Не удалось прочитать файл %q", file.Name()))
		return err
	}

	if err := ioutil.WriteFile(PathTo, buf, os.ModePerm); err != nil {
		PrintInfoLog(fmt.Sprintf("ОШИБКА! Не удалось записать файл %q", file.Name()))
		return err
	}

	return nil
}
