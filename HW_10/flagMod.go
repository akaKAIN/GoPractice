package main

import (
	"flag"
)

var (
	From   string
	To     string
	Offset int64
	Limit  int64
)

func InitFlag() {
	flag.StringVar(&From, "from", "", "Адрес копируемого файла")
	flag.StringVar(&To, "to", "", "Конечное расположение копируемого файла")
	flag.Int64Var(&Offset, "offset", 0, "Отступ в источнике.")
	flag.Int64Var(&Limit, "limit", 0, "Кол-во копируемых байт.")
}
