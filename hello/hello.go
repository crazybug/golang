package main

import (
         "fmt"
         "github.com/crazybug/stringutil"
)

func add(x, y int) int {
    return x + y
}

func mul(x, y int) int {
    return x*y
}

func main() {
    fmt.Println(stringutil.Reverse("!oG ,olleH"))
    fmt.Println(add(100,99))
    fmt.Println(mul(100,99))
}
