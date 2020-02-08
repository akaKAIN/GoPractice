package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//Инициализация меню-подсказки
	InitFlag()

	//Получаем список файлов в контролируемой директории
	FileList := GetFileList(DirFrom)

	//Проверка существования директории для конечных файлов. Если ее нет - создаем.
	PathToStorageDir := GetOrCreateDir(DirTo, ImageDir)

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
				err = DelFile(file, PathToStorageDir)
				if err != nil {
					log.Printf("DelError: %v file:%s\n", err, file.Name())
				}
			}

			//Информирование об успешном переносе файла
			PrintInfoLog(fmt.Sprintf("Перемещен: %s", file.Name()))
		}

	}

	fmt.Println(FileList, PathToStorageDir)

}

func PrintInfoLog(text string){
	if LOGGING{
		log.Println(text)
	}
}

func GetFileList(DirName string) []os.FileInfo {
	/*Возвращает список файлов указанной директории*/

	files, err := ioutil.ReadDir(DirName)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func GetOrCreateDir(dir, imageDir string) (fullPath string) {
	/*Проверяет существует ли указанная папка для хранения. Если нет - создает такую папку.
	Возвращает полный путь до папки*/

	absPath, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}
	fullPath = filepath.Join(absPath, imageDir)

	if _, err := os.Stat(fullPath); !os.IsExist(err) {
		if err := os.Mkdir(fullPath, os.ModePerm); err != nil {
			return
		}
		PrintInfoLog(fmt.Sprintf("Создана папка для хранения файлов по адресу: %v\n", fullPath))
	}
	return
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
		PrintInfoLog(fmt.Sprintf("ОШИБКА! Не удалось прочитать файл %q",file.Name()))
		return err
	}

	if err := ioutil.WriteFile(PathTo, []byte(buf), os.ModePerm); err != nil {
		PrintInfoLog(fmt.Sprintf("ОШИБКА! Не удалось записать файл %q",file.Name()))
		return err
	}

	return nil
}
