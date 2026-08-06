package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    args := os.Args[1:]
    if len(args) != 3 {
        return
    }

    a, err1 := strconv.ParseInt(args[0], 10, 64)
    b, err2 := strconv.ParseInt(args[2], 10, 64)
    if err1 != nil || err2 != nil {
        return
    }

    op := args[1]
    var result int64
    switch op {
    case "+":
        result = a + b
    case "-":
        result = a - b
    case "*":
        result = a * b
    case "/":
        if b == 0 {
            fmt.Println("No division by 0")
            return
        }
        result = a / b
    case "%":
        if b == 0 {
            fmt.Println("No modulo by 0")
            return
        }
        result = a % b
    default:
        return
    }

    // Overflow check (within int64 range)
    if result > (1<<63-1) || result < -(1<<63) {
        return
    }

    fmt.Println(result)
}