package main

import (
	"fmt"
	"net/http"
)

func getHeader(url string) {
	url = "http://" + url
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {

		fmt.Print(ColorYellow, k, ColorReset)
		fmt.Print(" : ")
		fmt.Println(v)
	}
}
