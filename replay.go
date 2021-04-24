package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var ignoreHeaders *string
var fileName *string
var client *http.Client = &http.Client{}
var request *http.Request

var hostName string
var httpVerb string
var httpPath string

var reachedBody bool = false

var requestContent string

func main() {
	ignoreHeaders = flag.String("i", "", "Ignore the headers which you dont want to add in request")
	fileName = flag.String("f", "", "File containing raw request")
	flag.Parse()

	//Compile different regex
	hostReg := regexp.MustCompile("(?i)Host: (.*)")                                            // regex to get Host value
	verbAndPath := regexp.MustCompile("(?i)(GET|POST|PUT|OPTIONS|DELETE|PATCH|UPDATE) (\\S*)") // regex to get http verb and path
	headerNameAndValue := regexp.MustCompile("(.*): (.*)")                                     // regex to get all headers

	// Read data from file passed via -f flag line by line
	count := 0
	f, errf := os.Open(*fileName)
	if errf != nil {
		fmt.Println(errf)
		os.Exit(1)
	} else {
		s := bufio.NewScanner(f)
		for s.Scan() {
			text := s.Text()
			if count == 0 {
				// if first line then capture verb and path from it
				match := verbAndPath.FindAllStringSubmatch(text, -1)
				httpVerb = match[0][1]
				httpPath = match[0][2]
			} else if count == 1 {
				// if second line capture host name from it
				match := hostReg.FindAllStringSubmatch(text, -1)
				hostName = match[0][1]

				// create a new http request via capture hostname verb and path
				request, _ = http.NewRequest(httpVerb, "https://"+hostName+httpPath, nil)
			} else {
				if reachedBody == true {
					requestContent = text
					break
				}

				if text == "" {
					reachedBody = true
					continue
				}
				// capture http headers from the file
				match := headerNameAndValue.FindAllStringSubmatch(text, -1)
				headerName := match[0][1]
				headerValue := match[0][2]

				//Check if header contains cookies if yes then add all the cookies in the reques
				if headerName == "Cookie" {
					listCookies := strings.Split(headerValue, " ")
					for _, eachCookie := range listCookies {
						cookieName := strings.Split(eachCookie, "=")[0]
						cookieValue := strings.Trim(strings.Split(eachCookie, "=")[1], ";")
						request.AddCookie(&http.Cookie{Name: cookieName, Value: cookieValue})
					}
				} else if headerName != *ignoreHeaders {
					// add http headers to the request
					request.Header.Add(headerName, headerValue)
				}
				// capture request body and add int Request = PENDING ====================
			}
			count++
		}
		//Add raw request Body in the request body

		finalRequest, _ := http.NewRequest(httpVerb, "https://"+hostName+httpPath, bytes.NewBuffer([]byte(requestContent)))
		finalRequest.Header = request.Header.Clone()
		for _, eachCookie := range request.Cookies() {
			finalRequest.AddCookie(eachCookie)
		}

		defer finalRequest.Body.Close()

		//make the final request
		resp, err := client.Do(finalRequest)

		defer resp.Body.Close()
		if err != nil {
			fmt.Println("Request failed", err)
		} else {
			fmt.Println(resp.Header)

			//add code to dump request body in correct format = PENDING ========================
			data, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(data))
		}

	}

}
