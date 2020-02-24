package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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
	NewFile  *os.FileInfo
}

func (s *Session) Copy() error {
	//Открываем файл на чтение
	file2, err := os.Open(s.PathFrom)
	CheckErr("OpenError", err)

	//Инициализируем слайс байтов для чтения
	lim := func() []byte {
		if s.Limit == 0 {
			return make([]byte, s.FileSize-s.Offset)
		} else {
			return make([]byte, s.Limit)
		}
	}()

	//Делаем смещение на указанной число байт
	position, err := file2.Seek(s.Offset, io.SeekStart)
	CheckErr("Ошибка расчета позиции Seek", err)

	//Читаем выбранный открезок
	readied, err := file2.ReadAt(lim, position)
	if err == io.EOF {
		msg := fmt.Sprintf("position: %d\nreadied: %d\nstr(r): %v\n\n", position, readied, string(lim))
		log.Println(msg)
	} else if err != nil {
		err = fmt.Errorf("Ошибка чтения: %s", err)
		return err
	}

	//Записываем считаный фрагмент в файл, предварительно его открыв
	CopeFile, err := os.OpenFile(To, os.O_RDWR, os.ModePerm)
	CheckErr("Ошибка открытия файла для копирования", err)
	defer func() {
		if err := CopeFile.Close(); err != nil {
			CheckErr("", err)
		}
	}()

	written, err := CopeFile.Write(lim)
	CheckErr("Ошибка записи буфера.", err)
	log.Printf("read: %d\nwrite:%d\n", readied, written)
	if readied != written {
		err = fmt.Errorf("Неверное количество данных скопировано.")
		CheckErr("", err)
	}

	return nil
}

func main() {
	flag.Parse() //Парсинг флогов.
	s, err := MakeSession()
	if err != nil {
		log.Println(err)
		CheckErr("MakeSession", err)
	}
	fmt.Printf("%+v\n", s)
	if err := s.Copy(); err != nil {
		CheckErr("Copy", err)
	}

	//file, err := os.Open(From)
	//
	//count := int(s.FileSize)
	//bar := pb.Full.Start(count)
	//
	//for i := 0; i < count; i++ {
	//	bar.Increment()
	//	//time.Sleep(time.Millisecond)
	//}
	//bar.Finish()

}
