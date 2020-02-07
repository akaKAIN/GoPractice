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

	//for _, file := range FileList {
	//
	//	//Проверяем файл на валидность и соответствие расширения
	//	if CheckFile(file) {
	//
	//		//Копируем файл в папку-хранилище
	//		err := CopyFile(file)
	//		if err != nil {
	//			log.Printf("CopyError: %v file:%s\n", err, file.Name())
	//		}
	//
	//		//Удаляем старую копию файла если файл успешно скопирован
	//		if err == nil {
	//			err = DelFile(file, PathToStorageDir)
	//			if err != nil {
	//				log.Printf("DelError: %v file:%s\n", err, file.Name())
	//			}
	//		}
	//
	//		//Информирование об успешном переносе файла
	//		if LOGGING && !err {
	//			log.Printf("Перемещен: %s", file.Name())
	//		}
	//	}
	//
	//}

fmt.Println(FileList, PathToStorageDir)

}

func GetFileList(DirName string) []os.FileInfo {

	files, err := ioutil.ReadDir(DirName)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func GetOrCreateDir(dir, imageDir string)(fullPath string){

	absPath, err := filepath.Abs(dir)
	if err !=nil{
		log.Fatal(err)
	}
	fullPath = filepath.Join(absPath, imageDir)

	if _, err := os.Stat(fullPath); !os.IsExist(err){
		if err := os.Mkdir(fullPath, os.ModePerm); err !=nil{
			return
		}
		log.Printf("Создана папка для хранения файлов по адресу: %v\n", fullPath)
	}
	return
}

func DelFile(file os.FileInfo, path string) error {
	PathToFile := filepath.Join(path, file.Name())
	if err := os.Remove(PathToFile); err != nil{
		return err
	}
	return nil
}
