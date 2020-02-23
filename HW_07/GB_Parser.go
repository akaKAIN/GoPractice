package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	Login(&wg)
	wg.Wait()
	for i := 0; i < 10000; i++ {
		go func() {
			i += 0
		}()
	}

}

func Login(wg *sync.WaitGroup) {
	const (
		urlLogin = "https://geekbrains.ru/login"
		urlParse = "https://geekbrains.ru/chapters/5414"
	)

	cookieJar, err := cookiejar.New(nil)
	CheckErr("Making New CookieJar", err)

	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           cookieJar,
		Timeout:       10 * time.Second,
	}

	resp, err := http.PostForm(urlLogin, url.Values{
		"user[email]":    {login},
		"user[password]": {password},
	})
	CheckErr("Sending PostForm", err)
	PrintResp(resp)

	GetSimpleTag(&client, urlParse, "a.lesson-header lesson-header_selected lesson-header_ended")
	CheckErr("Sending Get request", err)
	PrintResp(resp)
	wg.Done()
}

func PrintResp(r *http.Response) {
	fmt.Printf("Status: %s\n, header: %s\n\n", r.Status, r.Header)

}
