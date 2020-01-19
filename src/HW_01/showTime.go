package main

import (
	"fmt"
	"time"
)

func main() {

	//Mon Jan 2 15:04:05 MST 2006.
	fmt.Println(time.Now().Format("15:04"))
}
