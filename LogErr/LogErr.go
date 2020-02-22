package LogErr

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	inFile, _ = os.OpenFile("error.log", os.O_CREATE|os.O_APPEND, os.ModePerm)
	LogThis = log.New(inFile, fmt.Sprint(time.Now().Format("02.01.2006 15:04:05 ")), 0)
)

func CheckErr(err error){
	if err != nil {
		LogThis.Println(err)
	}
}