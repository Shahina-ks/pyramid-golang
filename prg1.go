package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter the number of rows: ")
    fmt.Scan(&n)

    c := 'w' //to start with w
    for i := 0; i < n; i++ {
        for j := 0; j <= i; j++ {
            fmt.Printf("%c", c)
        }
        fmt.Println()
        c++
    }
}


