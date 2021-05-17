package main

import (
    "fmt"
)

func main() {
    var x float32
    fmt.Println("Enter a floating point number : ")
    fmt.Scan(&x)
    fmt.Printf("Truncation is the process : %d", int16(x))
}