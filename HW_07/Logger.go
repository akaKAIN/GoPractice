package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	inFile, _ = os.OpenFile("error.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModePerm)
	LogThis   = log.New(inFile, fmt.Sprint(time.Now().Format("02.01.2006 15:04:05 ")), 0)
)

func CheckErr(prefix string, err error) {
	if err != nil {
		LogThis.Printf("%s: %v", prefix, err)
	}
}
