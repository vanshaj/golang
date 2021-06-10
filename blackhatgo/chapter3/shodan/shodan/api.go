package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIInfo struct {
	Member       bool   `json:"member"`
	Credits      int    `json:"credits"`
	Display_name string `json:"display_name"`
	Created      string `json:"created"`
}

func (c *Client) APIInfo() (*APIInfo, error) {
	endpoint := fmt.Sprintf("%s/account/profile?key=%s", BASE_URL, c.apiKey)
	fmt.Println(endpoint)
	httpC := &http.Client{}
	res, err := httpC.Get(endpoint)
	if err != nil {
		fmt.Println("Unable to make get calL ", err)
		return nil, err
	}
	defer res.Body.Close()
	var data APIInfo
	err2 := json.NewDecoder(res.Body).Decode(&data)
	return &data, err2
}
