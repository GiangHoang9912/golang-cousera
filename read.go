package main

import (
    "fmt"
	"os"
	"bufio"
	"io"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	fmt.Println("Enter File Path : ")
	var file_path string

	fmt.Scan(&file_path)

	var slides []Person

    file, err := os.Open(file_path)
	if err != nil {
        fmt.Println(err)
    }

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		list_data := strings.Split(string(line), " ")
		p := Person{
			list_data[0],
			list_data[1]}
		slides = append(slides, p)
	}

	fmt.Println(slides)
}
