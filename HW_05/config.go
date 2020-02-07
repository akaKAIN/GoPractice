package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

var (
	DirFrom string
	DirTo	string
	TIMER   int
	LOGGING bool
)

const ImageDir = "Images"

func GetBaseDir() string {
	way, _ := filepath.Abs(".")
	return filepath.Base(way)
}

func InitFlag() {
	flag.IntVar(&TIMER, "timer", 90, "Интервал проверки (сек).")
	flag.StringVar(&DirFrom, "dir_from", GetBaseDir(), "Директория для поиска файлов")
	flag.StringVar(&DirTo, "dir_to", GetBaseDir(), "Директория для сохранения файлов")
	flag.BoolVar(&LOGGING, "log", false, "Нужно ли выводить логи копирования файлов (Да=1/true; Нет=0/false).")
	flag.Parse()
	tailPtr := flag.Args()
	fmt.Printf("directory: %s\n", DirFrom)
	fmt.Printf("directory: %s\n", DirTo)
	fmt.Printf("timer: %d\n", TIMER)
	fmt.Printf("Logging: %t\n", LOGGING)
	fmt.Printf("tail: %v\n", tailPtr)
}