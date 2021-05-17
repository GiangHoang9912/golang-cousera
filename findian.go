package main

import (
    "fmt"
    "strings"
)

func main() {
    var s string
    fmt.Println("Enter a string : ")
    fmt.Scan(&s)
    var sub_i rune = 'i'
    var sub_n rune = 'n'
    var sub_a string = "a"

    i := strings.IndexRune(strings.ToLower(s), sub_i)
    n := strings.IndexRune(strings.ToLower(s), sub_n)
    a := strings.Contains(strings.ToLower(s), sub_a)
    
    if a && i == 0 && n == len(s) - 1 {
        fmt.Println("result : Found!")
        return
    }

    fmt.Println("result: Not Found!")
}