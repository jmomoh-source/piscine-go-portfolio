# Quest11 — listfind

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **searching in linked lists** in Go.  
The task: write a function `ListFind` that returns the address of the `Data` field of the first node in the list that matches a reference value, using a comparator function.

Rules:
- Use the provided `NodeL` and `List` structures.
- Use the comparator function `CompStr` to check equality.
- If a match is found, return the pointer to the `Data` field.
- If no match is found, return `nil`.

## Instructions
- File to submit: `listfind.go`
- Expected structures and function signatures:
```go
type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func CompStr(a, b interface{}) bool {
    return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{}
```

## Implementation
`listfind.go`:
```go
package piscine

type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func CompStr(a, b interface{}) bool {
    return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
    current := l.Head
    for current != nil {
        if comp(current.Data, ref) {
            return &current.Data
        }
        current = current.Next
    }
    return nil
}
```

### Explanation
- Traverse the list starting from `Head`.
- For each node, compare its `Data` with `ref` using `comp`.
- If equal, return the address of the `Data` field.
- If traversal ends without a match, return `nil`.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    link := &piscine.List{}

    piscine.ListPushBack(link, "hello")
    piscine.ListPushBack(link, "hello1")
    piscine.ListPushBack(link, "hello2")
    piscine.ListPushBack(link, "hello3")

    found := piscine.ListFind(link, interface{}("hello2"), piscine.CompStr)

    fmt.Println(found)   // prints memory address
    fmt.Println(*found)  // prints value
}
```

Output:
```text
0xc42000a0a0
hello2
```
⚠️ Note: The memory address will differ between executions.

## Standard Library Equivalent
Go’s `container/list` package does not provide a built‑in search.  
You can implement it manually:
```go
import (
    "container/list"
    "fmt"
)

func Find(l *list.List, ref interface{}) *interface{} {
    for e := l.Front(); e != nil; e = e.Next() {
        if e.Value == ref {
            return &e.Value
        }
    }
    return nil
}

func main() {
    l := list.New()
    l.PushBack("hello")
    l.PushBack("hello1")
    l.PushBack("hello2")

    found := Find(l, "hello2")
    fmt.Println(found)   // memory address
    fmt.Println(*found)  // hello2
}
```

## Skills Practiced
- Structs and pointers
- Linked list traversal
- Comparator functions
- Returning pointers to data fields

## Notes
- This exercise demonstrates how to search for values in a linked list.
- Unlike slices, linked lists require traversal for search operations.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Effective Go — Data structures (go.dev in Bing) 