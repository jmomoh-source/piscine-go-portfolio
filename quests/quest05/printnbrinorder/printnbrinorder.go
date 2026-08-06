package main

import "fmt"

func main() {
    PrintNbrInOrder(321)
    PrintNbrInOrder(0)
    PrintNbrInOrder(321)
}

func PrintNbrInOrder(n int) {
    if n == 0 {
        fmt.Print('0')
        return
    }

    digits := []int{}
    for n > 0 {
        digits = append(digits, n%10)
        n /= 10
    }

    // Simple insertion sort
    for i := 1; i < len(digits); i++ {
        key := digits[i]
        j := i - 1
        for j >= 0 && digits[j] > key {
            digits[j+1] = digits[j]
            j--
        }
        digits[j+1] = key
    }

    for _, d := range digits {
        fmt.Print(rune(d + '0'))
    }
}