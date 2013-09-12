package main

import (
	"fmt"
	//"io"
	"net/http"
	"net/url"
	"os"
	//"regexp"
)

func main() {
	account := os.Args[1]
	passwd := os.Args[2]

	var body = url.Values{"DDDDD": {account}, "upass": {passwd}, "R1": {"0"}}

	URL := "http://10.3.8.211"

	//var client http.Client
	res, err := http.PostForm(URL, body)
	//	res, err := http.Get(URL)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var buf = make([]byte, 200)

	res.Body.Read(buf)
	fmt.Println(string(buf))
}
