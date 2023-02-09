package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api-football-v1.p.rapidapi.com/v3/timezone"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "0297eb7b3bmsh4c7bb19a9822296p1efde8jsnfb75673da3af")
	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")

	fmt.Println("req:", req)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	//

}
