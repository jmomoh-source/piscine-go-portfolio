package main

import (
    "fmt"
    "sort"
)

func main() {
    middle := Abort(2, 3, 8, 5, 7)
    fmt.Println(middle) // 5
}

func Abort(a, b, c, d, e int) int {
    nums := []int{a, b, c, d, e}
    sort.Ints(nums)
    return nums[2] // median is the 3rd element (index 2)
}
