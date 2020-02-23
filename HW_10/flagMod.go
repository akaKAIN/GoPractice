package main

import (
	"flag"
)

var (
	From   string
	To     string
	Offset int
	Limit int
)

func InitFlag() {
	flag.StringVar(&From, "from", "", "Адрес копируемого файла")
	flag.StringVar(&To, "to", "", "Конечное расположение копируемого файла")
	flag.IntVar(&Offset, "offset", 0, "Отступ в источнике.")
	flag.IntVar(&Limit, "limit", 0, "Кол-во копируемых байт.")
}
