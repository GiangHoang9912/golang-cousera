package main

import (
    "fmt"
    "strings"
    "os"
    "strconv"
)

func main() {
    var arr [3]int

    var slides []int = arr[0:3]

    x := ""

    for i:=0; i < len(slides); i++{
        fmt.Println("Enter Number : ")
        fmt.Scan(&x)
        if strings.ToLower(x) == "x" {
            os.Exit(0)
        }
        val, err := strconv.Atoi(x)
        slides[i] = val

        if err != nil {
            fmt.Println(err)
            return
        }
    }
    sorted := sort(slides)

    fmt.Println(sorted)

    for {
        fmt.Println("Enter Number : ")
        fmt.Scan(&x)
        if strings.ToLower(x) == "x" {
            os.Exit(0)
        }

        i, err := strconv.Atoi(x)

        if err != nil {
            continue
        }

        slides = append(slides, i)

        sorted := sort(slides)
        fmt.Println(sorted)
    }
}

func sort(slides []int) []int {
    var temp int
    for i := 0;i < len(slides) - 1; i++ {
        for j := 0; j < len(slides) -i -1;j++ {
            if (slides[j] > slides[j+1]) {
                temp = slides[j]
                slides[j] = slides[j+1]
                slides[j+1] = temp
            }
        }
    }
    return slides
} 