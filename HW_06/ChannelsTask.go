package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task struct {
	Func     func(chan<- error, <- chan string)
	Count    int
	ErrLimit int
}

func main() {
	NewTask := Task{
		Func:     PrintNum,
		Count:    10,
		ErrLimit: 2,
	}


	HandlerFunc(NewTask)

}

func HandlerFunc(task Task) {
	count := 0
	ch := make(chan error, task.ErrLimit)
	trigger := make(chan string)
	for i := 0; i < task.Count; i++ {
		go task.Func(ch, trigger)
	}
	for {
		_, ok := <- ch
		if ok {
			count++
		}
		if count > task.ErrLimit{
			trigger <- "stop"
		}
	}
	fmt.Println("Случилось ошибок:", count)
}

func PrintNum(out chan<- error, in <- chan string) {
	count := rand.Intn(10)
	pause := rand.Intn(10)
	err := fmt.Errorf("Error of printing.\n")
	if count > pause && pause != 0 {

		for i := 0; i < count; i++ {
			fmt.Println("Print count num:", count)
			time.Sleep(time.Second)
			for _, ok := range in{

			}
		}
	} else {
		out <- err
	}

}

