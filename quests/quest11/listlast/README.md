# Quest11 — listlast

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **linked list traversal to the last node** in Go.  
The task: write a function `ListLast` that returns the `Data` of the last element in a linked list.

Rules:
- Use the provided `NodeL` and `List` structures.
- If the list is empty, return `nil`.
- Otherwise, return the `Data` of the last node.

## Instructions
- File to submit: `listlast.go`
- Expected structures and function signature:
```go
type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func ListLast(l *List) interface{}
```

## Implementation
`listlast.go`:
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

func ListLast(l *List) interface{} {
    if l.Head == nil {
        return nil
    }
    return l.Tail.Data
}
```

### Explanation
- If the list is empty (`Head == nil`), return `nil`.
- Otherwise, return the `Data` stored in the `Tail` node.
- Since the `Tail` pointer is maintained during insertions, this is efficient (O(1)).

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
    link2 := &piscine.List{}

    piscine.ListPushBack(link, "three")
    piscine.ListPushBack(link, 3)
    piscine.ListPushBack(link, "1")

    fmt.Println(piscine.ListLast(link))
    fmt.Println(piscine.ListLast(link2))
}
```

Output:
```text
1
<nil>
```

## Standard Library Equivalent
Go’s `container/list` package provides a doubly linked list with a built‑in `Back()` method:
```go
import (
    "container/list"
    "fmt"
)

func main() {
    l := list.New()
    l.PushBack("three")
    l.PushBack(3)
    l.PushBack("1")

    fmt.Println(l.Back().Value) // 1
    empty := list.New()
    fmt.Println(empty.Back())   // <nil>
}
```
⚠️ Note: `container/list.Back()` is concise and idiomatic.  
Your Piscine solution demonstrates how to manually retrieve the last node in a custom singly linked list.

## Skills Practiced
- Structs and pointers
- Linked list traversal
- Efficient use of head and tail references
- Handling empty list cases

## Notes
- This exercise demonstrates how to access the last node in a linked list.
- For production code, prefer `container/list` unless you need a custom implementation.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Effective Go — Data structures (go.dev in Bing) 