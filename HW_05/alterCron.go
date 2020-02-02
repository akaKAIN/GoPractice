package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	var count int
	var path = "."

	if len(os.Args) == 2 {
		checkPath := os.Args[1]
		if _, err := os.Stat(checkPath); !os.IsNotExist(err) {
			path = checkPath
		}
	}

	for {
		count++
		fmt.Printf("\nStep № %d\n", count)
		MoveImages(path)
		time.Sleep(60 * time.Second)

	}
}

func MoveImages(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}

}
