package main

import "fmt"

func main() {
    PrintCombN(1)
    PrintCombN(3)
    PrintCombN(9)
}

func PrintCombN(n int) {
    if n <= 0 || n >= 10 {
        return
    }
    comb := make([]int, n)
    generateComb(0, 0, n, comb)
    fmt.Println()
}

func generateComb(pos, start, n int, comb []int) {
    if pos == n {
        for i := 0; i < n; i++ {
            fmt.Printf("%d", comb[i])
        }
        // Check if this is the last combination
        if comb[0] != 10-n {
            fmt.Print(", ")
        }
        return
    }
    for d := start; d <= 9; d++ {
        comb[pos] = d
        generateComb(pos+1, d+1, n, comb)
    }
}