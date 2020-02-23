package main

import (
	"flag"
	"github.com/cheggaaa/pb/v3"
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
	Offset   int
	Limit    int
}

func main() {
	flag.Parse() //Парсинг флогов.


	//file, err := os.Open(From)

	count := 10000
	// start bar from 'full' template
	bar := pb.Full.Start(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()

}
