package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var count int
	var path = "."

	if len(os.Args) == 2 {
		checkPath := os.Args[1]
		if _, err := os.Stat(checkPath); !os.IsNotExist(err) {
			path = checkPath
		}
	}

	for {
		count++
		fmt.Printf("\nStep № %d\n", count)
		MoveImages(path)
		time.Sleep(60 * time.Second)

	}
}

func MoveImages(path string) {
	var formats = []string{".jpeg", ".jpg"}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		for _, format := range formats {
			if strings.Contains(file.Name(), format) {
				//TODO: replace func
				err := replace(file)
				if err != nil {
					log.Printf("Ошибка перемещения файла: %q", file)
				}
			}
		}
	}

}

func replace(file os.FileInfo) error {
	return nil
}
