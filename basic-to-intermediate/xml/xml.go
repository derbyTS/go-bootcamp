package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email,omitempty"`
	Address Address  `xml:"address"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {
	person := Person{
		Name: "John",
		Age:  32,
		Address: Address{
			City:  "Manila",
			State: "Metro Manila",
		},
	}

	personData, err := xml.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(personData))

	fmt.Println(string("------------------------"))

	personData, err = xml.MarshalIndent(person, "-->", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(personData))

	rawData := `<person><name>John</name><age>32</age><address><city>Manila</city><state>Metro Manila</state></address></person>`

	var personXml Person
	err = xml.Unmarshal([]byte(rawData), &personXml)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(personXml)
	fmt.Println("Local string: ", personXml.XMLName.Local)
	fmt.Println("Namespace: ", personXml.XMLName.Space)
	book := Book{
		ISBN:      "123-123-123-123",
		Title:     "Gogogogogo",
		Author:    "John",
		Pages:     200,
		Publisher: "Publish Company",
	}

	bookDataAttr, err := xml.MarshalIndent(book, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bookDataAttr))
}

type Book struct {
	XMLName   xml.Name `xml:"book"`
	ISBN      string   `xml:"isbn,attr"`
	Title     string   `xml:"title,attr"`
	Author    string   `xml:"author,attr"`
	Pages     int      `xml:"pages"`
	Publisher string   `xml:"publisher,attr"`
}
