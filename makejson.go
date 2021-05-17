package main

import (
    "fmt"
	"encoding/json"
)

func main() {
	var name string
	fmt.Println("Enter Person Name : ")
	fmt.Scan(&name)

	var address string
	fmt.Println("Enter Address Name : ")
	fmt.Scan(&address)

	m := make(map[string]string)

	m["name"] = name
	m["address"] = address

	fmt.Println(m)

	jsonString, err := json.Marshal(m)
	fmt.Println(err)
	fmt.Println(jsonString)
}
