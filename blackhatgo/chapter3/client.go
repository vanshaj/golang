package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var client *http.Client = &http.Client{}
	baseurl := "https://reqres.in/api/"
	endpoint := "users?page=2"

	req, err := http.NewRequest("GET", baseurl+endpoint, nil)
	if err != nil {
		fmt.Println("Unable to create connection to ", baseurl, " because of ", err)
		return
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Unable to send request ", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body) // resp.Body is io.ReadCloser and ReadAll return []bytes
	if err != nil {
		fmt.Println("Unable to read body")
		return
	}
	fmt.Println(string(body))
}
