package main

import (
	"fmt"

	"github.com/blackhatgo/chapter3/shodan/shodan"
)

func main() {
	shodanKey := ""
	client := shodan.New(shodanKey)
	res, err := client.APIInfo()
	if err != nil {
		fmt.Println("Unable to get APIINfo ", err)
		return
	}
	fmt.Println(res.Display_name)
	fmt.Println(res.Credits)
}
