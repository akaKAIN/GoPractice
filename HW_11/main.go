package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func signalHandler(c <- chan os.Signal){
	s := <- c
	fmt.Println("Get signal:", s)
	os.Exit(1)
}

func main(){
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	go signalHandler(c)
	ShowEnv()
}
