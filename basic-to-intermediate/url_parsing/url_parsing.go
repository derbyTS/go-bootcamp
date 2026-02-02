package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawUrl := "https://example.com:8080/path?query=param#fragment"
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("RawQuery:", parsedUrl.RawQuery)
	fmt.Println("Fragment:", parsedUrl.Fragment)

	fmt.Println("------------------------------------")

	rawUrl1 := "https://example.com:8080/path?name=Gracie&age=34"
	parsedUrl1, err := url.Parse(rawUrl1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Scheme:", parsedUrl1.Scheme)
	fmt.Println("Host:", parsedUrl1.Host)
	fmt.Println("Port:", parsedUrl1.Port())
	fmt.Println("Path:", parsedUrl1.Path)
	fmt.Println("RawQuery:", parsedUrl1.RawQuery)
	fmt.Println("Fragment:", parsedUrl1.Fragment)
	queryParams := parsedUrl1.Query()
	fmt.Println("Query Params:", queryParams)
	fmt.Println("Query Params name:", queryParams.Get("name"))
	fmt.Println("Query Params age:", queryParams.Get("age"))

	fmt.Println("-----------------Build URL-------------------")
	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseUrl.Query()
	query.Set("name", "Gracie")
	baseUrl.RawQuery = query.Encode() // If you not encode it, it wont add to url

	fmt.Println("Built URL: ", baseUrl.String())

	values := url.Values{}

	values.Add("name", "Gracie")
	values.Add("age", "34")
	values.Add("city", "bacoor")
	encoded := values.Encode()
	fmt.Println(encoded)

	baseUrl1 := "https://example.com"
	fullUrl := baseUrl1 + "?" + encoded

	fmt.Println(fullUrl)
}
