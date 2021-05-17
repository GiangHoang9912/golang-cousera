package main

import (
    "fmt"
	"encoding/json"
)

type Person struct {
	name string
	address string
	phone string
}

func main() {
    x := Person {
		"giang",
		"Thai Nguyen",
		"0819169868"}

	bar, err := json.Marshal(x)

	fmt.Println(bar)
	fmt.Println(err)
}
