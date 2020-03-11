package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

//go:generate echo "generator"

func ShowEnv() {
	var env []string
	env = os.Environ()
	fmt.Println(env[:5])
	fmt.Println(os.LookupEnv("LANG"))

	cmd := exec.Command("ls", "-lsa")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()

	out, err := exec.Command("ls").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", string(out))

	for {
		fmt.Println("LOOP")
		time.Sleep(time.Second * 2)
	}

}
