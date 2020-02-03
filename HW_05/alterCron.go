package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)



func main() {
	var count int
	var workDir = "."

	if len(os.Args) == 2 {
		checkPath := os.Args[1]
		if _, err := os.Stat(checkPath); !os.IsNotExist(err) {
			workDir = checkPath
		}
	}

	for {
		count++
		fmt.Printf("\nStep № %d\n", count)
		MoveImages(workDir)
		time.Sleep(60 * time.Second)

	}
}

func MoveImages(path string) {
	var formats = []string{".jpeg", ".jpg"}

	AbsPath, _ := filepath.Abs(ImageDir)

	err := os.MkdirAll(ImageDir, os.ModePerm)
	if err == nil {
		fmt.Printf("Создана папка для хранения картинок по адресу:\n%s\n\n", AbsPath)
	}


	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		for _, format := range formats {
			if filepath.Ext(file.Name()) == format {
				//TODO: replace func
				err := replace(file)
				if err != nil {
					log.Printf("Ошибка перемещения файла: %q", file)
				}
				continue
			}
		}
	}

}

func replace(file os.FileInfo) error {
	AbsPath, _ := filepath.Abs(ImageDir)
	FilePath := path.Join(AbsPath, file.Name())

	body, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(FilePath, body, 775); err != nil {
		return err
	}

	fmt.Printf("Файл %s скопирован.", file.Name())
	return nil
}
