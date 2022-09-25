package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

//get URL and return string array
func getURL(params ...string) ([]string, http.Header, string) {

	url := params[0]
	token := params[1]
	custHeader := params[2]

	//create http client
	client := &http.Client{}

	//fetch http.get
	req, err := http.NewRequest("GET", url, nil)

	if len(token) > 0 {
		req.Header.Set("Authorization", token)
	}

	if len(custHeader) > 0 {
		fmt.Println(custHeader)
		strs := strings.Split(custHeader, ":")
		req.Header.Set(strs[0], strs[1])
	}

	fmt.Println(req.Header)
	//resp, err := http.Get(url)
	resp, err := client.Do(req)

	//check for error
	if err != nil {
		log.Fatal(err)
	}

	//close after use
	defer resp.Body.Close()
	//defer req.Body.Close()

	//declare string array to store ascii chars
	var charBody []string

	//read Body of http response
	body, err := io.ReadAll(resp.Body)
	//body, err := io.ReadAll(req.Body)

	//fmt.Println(err)
	for _, v := range body {
		cha := string(v)
		charBody = append(charBody, cha)
	}
	return charBody, resp.Header, resp.Status
}

func main() {

	type CustomHeader struct {
		Name  string
		Value string
	}

	//parse CLI args; (char, default, description)
	argUrl := flag.String("u", "", "URL")
	argHeader_only := flag.Bool("h", false, "only show header response")
	argToken := flag.String("t", "", "Auth token")
	argCustHeader := flag.String("c", "", "Header_Name: Value")
	flag.Parse()

	//dereference arg pointer to var
	url := *argUrl
	headOnly := *argHeader_only
	authToken := *argToken
	custHeader := *argCustHeader

	//sanitize user URL entry
	if strings.Contains(url, "://") == false {
		url = "https://" + url
	}

	//make request...

	//display request on stdout
	fmt.Println("Requesting: " + url)

	//call func and return string array
	charArray, respHeader, respStatus := getURL(url, authToken, custHeader)

	//display results...
	fmt.Println("Status Code: " + respStatus + "\n")

	//print to stdout
	if headOnly == true {
		fmt.Println("***** Header *****")
		for i, v := range respHeader {
			fmt.Println(i, " ", v)
		}

	} else {
		fmt.Println("***** Header *****")
		for i, v := range respHeader {
			fmt.Println(i, " ", v)
		}
		fmt.Println("")
		fmt.Println("***** Response ***** ")

		for _, v := range charArray {
			fmt.Print(v)
		}
	}
	fmt.Println()
}
