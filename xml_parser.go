package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Users struct {
	XMLName xml.Name `xml:"users"`
	Emp     []User   `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name"`
	Type    string   `xml:"type,attr"`
	Social  Social   `xml:"social"`
}

type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func main() {
	xmlFile, err := os.Open("users.xml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.xml File")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var users Users
	xml.Unmarshal(byteValue, &users)

	fmt.Println(users)
	for i := 0; i < len(users.Emp); i++ {
		fmt.Println("User Type : " + users.Emp[i].Type)
		fmt.Println("User Name : " + users.Emp[i].Name)
		fmt.Println("Facebook Link : " + users.Emp[i].Social.Facebook)
		fmt.Println("Twitter Link : " + users.Emp[i].Social.Twitter)
		fmt.Println("Youtube Link : " + users.Emp[i].Social.Youtube)
	}
}
