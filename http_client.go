package main

import (
	"net/http"
	"strings"
)

const url = "http://localhost:8080/"

func main() {
	//link := "https://github.com/popov-andrei"
	//link := "https://vk.com/coding21"
	link := "https://www.apple.com/iphone-14-pro/"

	resp, _ := http.Post(url, "text/plain", strings.NewReader(link))
	defer resp.Body.Close()
}
