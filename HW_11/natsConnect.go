package main

import (
	"encoding/json"
	"fmt"
	nats "github.com/nats-io/nats.go"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Town string `json:"town"`
}

type Community struct {
	All []Person    `json:"all"`
}

var Res Community

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	//go SubSc("new", nc)
	SubSc, err := nc.Subscribe("new", MySub)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = SubSc.Unsubscribe()
		if err != nil {
			log.Fatal(err)
		}
		if SaveResult(){
			fmt.Println("All was done very well.")
		}
	}()

	for i := 20; i < 60; i++ {
		var person = Person{"Ivan", i, "Moscow"}

		data, err := json.Marshal(person)
		if err != nil {
			log.Fatal(err)
		}

		if err = nc.Publish("new", data); err != nil {
			log.Fatal(err)
		}

	}
	time.Sleep(time.Second * 1)
}

func SubSc(key string, nc *nats.Conn) {
	sub, err := nc.SubscribeSync(key)
	if err != nil {
		log.Fatal(err)
	}

	timeOut := time.Second * 10
	msg, err := sub.NextMsg(timeOut)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(msg.Data))
}

func MySub(m *nats.Msg) {
	var p = Person{}
	if err := json.Unmarshal(m.Data, &p); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Get:\n%#v\n", p)
	Res.All = append(Res.All, p)
}

func SaveResult() bool {
	data, err := json.Marshal(Res)
	if err != nil {
		log.Println(err)
		return false
	}
	if err = ioutil.WriteFile("result_persons.json", data, os.ModePerm); err != nil {
		log.Println(err)
		return false
	}
	return true
}
