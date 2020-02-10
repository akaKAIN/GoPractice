package main

import (
	"flag"
	"log"
	"path/filepath"
)

var (
	DirFrom string
	DirTo	string
	TIMER   int
	LOGGING bool
)

var ExtArr = []string{".jpeg", ".jpg", ".png"}

const ImageDir = "Images"

func GetBaseDir() string {
	way, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}
	return way
}

func GetStorageDir() string {
	return filepath.Join(DirFrom, ImageDir)
}

func InitFlag() {
	flag.StringVar(&DirFrom, "dir_from", GetBaseDir(), "Директория для поиска файлов")
	flag.StringVar(&DirTo, "dir_to", GetStorageDir(), "Директория для сохранения файлов")
	flag.BoolVar(&LOGGING, "log", false, "Нужно ли выводить логи копирования файлов (Да=1/true; Нет=0/false).")
	flag.IntVar(&TIMER, "timer", 90, "Интервал проверки (сек).")
	flag.Parse()
	//tailPtr := flag.Args()
	//fmt.Printf("directory: %s\n", DirFrom)
	//fmt.Printf("directory: %s\n", DirTo)
	//fmt.Printf("timer: %d\n", TIMER)
	//fmt.Printf("Logging: %t\n", LOGGING)
	//fmt.Printf("tail: %v\n", tailPtr)
}