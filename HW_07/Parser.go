package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func GetSimpleTag(client *http.Client, url, selector string) {
	resp, err := client.Get(url)
	defer func(){
		err = resp.Body.Close()
		CheckErr("Closing resp.Body", err)
	}()
	CheckErr("GetSimpleGet - Fatal", err)
	PrintResp(resp)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	CheckErr("SimpleParse - Fatal", err)
	text, ok := doc.Find(selector).Attr("href")
	if !ok{
		err = fmt.Errorf("No data in tag: %q by url: %s\n", selector, url)
		CheckErr("Parse error", err)
	}else{
		fmt.Println(text)
	}
}
