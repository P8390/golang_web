package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Emp []User `json:"users"`
}
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json File")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValue, &users)

	fmt.Println(users)
	for i := 0; i < len(users.Emp); i++ {
		fmt.Println("User Type : " + users.Emp[i].Type)
		fmt.Println("User Name : " + users.Emp[i].Name)
		fmt.Println("User Age : " + strconv.Itoa(users.Emp[i].Age))
		fmt.Println("Facebook Link : " + users.Emp[i].Social.Facebook)
		fmt.Println("Twitter Link : " + users.Emp[i].Social.Twitter)
	}
}
