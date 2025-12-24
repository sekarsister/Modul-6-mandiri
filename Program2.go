package main

import "fmt"

func main() {
    var a int
    fmt.Scan(&a)
    
    if a > 0 {
        for i := a; i >= 1; i-- {
            for j := 1; j <= i; j++ {
                fmt.Print(j, " ")
            }
            fmt.Println()
        }
    } else if a == 0 {
        fmt.Println("-")
    }
}
