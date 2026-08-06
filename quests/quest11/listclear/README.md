# Quest11 — listclear

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **linked list deletion** in Go.  
The task: write a function `ListClear` that deletes all nodes from a linked list.

Rules:
- Use the provided `NodeL` and `List` structures.
- After clearing, both `Head` and `Tail` must be `nil`.
- Traversal is not required; simply reset the list pointers.

## Instructions
- File to submit: `listclear.go`
- Expected function signature:
```go
func ListClear(l *List)
```

## Implementation
`listclear.go`:
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

func ListClear(l *List) {
    l.Head = nil
    l.Tail = nil
}
```

### Explanation
- Assign `Head` and `Tail` to `nil`.
- This effectively clears the list, since no nodes are reachable anymore.
- Go’s garbage collector will free the memory of the orphaned nodes.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

type List = piscine.List
type Node = piscine.NodeL

func PrintList(l *List) {
    link := l.Head
    for link != nil {
        fmt.Print(link.Data, " -> ")
        link = link.Next
    }
    fmt.Println(nil)
}

func main() {
    link := &List{}

    piscine.ListPushBack(link, "I")
    piscine.ListPushBack(link, 1)
    piscine.ListPushBack(link, "something")
    piscine.ListPushBack(link, 2)

    fmt.Println("------list------")
    PrintList(link)
    piscine.ListClear(link)
    fmt.Println("------updated list------")
    PrintList(link)
}
```

Output:
```text
------list------
I -> 1 -> something -> 2 -> <nil>
------updated list------
<nil>
```

## Standard Library Equivalent
Go’s `container/list` package provides a doubly linked list with an `Init()` method that clears the list:
```go
import (
    "container/list"
    "fmt"
)

func main() {
    l := list.New()
    l.PushBack("I")
    l.PushBack(1)
    l.PushBack("something")
    l.PushBack(2)

    fmt.Println(l.Len()) // 4
    l.Init()             // clears the list
    fmt.Println(l.Len()) // 0
}
```
⚠️ Note: `Init()` resets the list to empty.  
Your Piscine solution demonstrates how to manually clear a custom singly linked list.

## Skills Practiced
- Structs and pointers
- Linked list memory management
- Resetting head and tail references
- Understanding garbage collection in Go

## Notes
- This exercise demonstrates how to clear a linked list efficiently.
- In Go, explicit memory freeing is not required; garbage collection handles it.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Effective Go — Data structures (go.dev in Bing)