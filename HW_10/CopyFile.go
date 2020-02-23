package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"log"
	"time"
)

/*
Копирование файлов
Цель: Реализовать утилиту копирования файлов Утилита должна принимать следующие аргументы
* файл источник (From)
* файл копия (To)
* Отступ в источнике (Offset), по умолчанию - 0
* Количество копируемых байт (Limit), по умолчанию - весь файл из From.
Выводить в консоль прогресс копирования в %,
например с помощью github.com/cheggaaa/pb Программа может НЕ обрабатывать файлы, у которых не известна длинна
(например /dev/urandom).
*/
func init() {
	InitFlag()
}

type Session struct {
	PathFrom string
	PathTo   string
	FileName string
	FileSize int64
	Offset   int64
	Limit    int64
	NewFile *bufio.Writer
}

func main() {
	flag.Parse() //Парсинг флогов.
	s, err := MakeSession()
	if err != nil {
		log.Println(err)
		CheckErr("MakeSession", err)
	}
	fmt.Printf("%+v\n", s)


	//file, err := os.Open(From)

	count := 2000
	bar := pb.Full.Start(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()

}
