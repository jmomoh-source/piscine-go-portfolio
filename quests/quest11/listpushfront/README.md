# Quest11 — listpushfront

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **linked list insertion at the front** in Go.  
The task: write a function `ListPushFront` that inserts a new node at the beginning of a linked list.

Rules:
- Use the provided `NodeL` and `List` structures.
- Update both `Head` and `Tail` pointers correctly.
- Handle the case where the list is initially empty.

## Instructions
- File to submit: `listpushfront.go`
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

func ListPushFront(l *List, data interface{})
```

## Implementation
`listpushfront.go`:
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

func ListPushFront(l *List, data interface{}) {
    newNode := &NodeL{Data: data}
    if l.Head == nil {
        l.Head = newNode
        l.Tail = newNode
    } else {
        newNode.Next = l.Head
        l.Head = newNode
    }
}
```

### Explanation
- Create a new node with the given data.
- If the list is empty, set both `Head` and `Tail` to the new node.
- Otherwise, link the new node to the current `Head` and update `Head`.

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

    piscine.ListPushFront(link, "Hello")
    piscine.ListPushFront(link, "man")
    piscine.ListPushFront(link, "how are you")

    it := link.Head
    for it != nil {
        fmt.Print(it.Data, " ")
        it = it.Next
    }
    fmt.Println()
}
```

Output:
```text
how are you man Hello
```

## Standard Library Equivalent
Go does not have a built‑in singly linked list type, but the `container/list` package provides a doubly linked list:
```go
import (
    "container/list"
    "fmt"
)

func main() {
    l := list.New()
    l.PushFront("Hello")
    l.PushFront("man")
    l.PushFront("how are you")

    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Print(e.Value, " ")
    }
    fmt.Println()
}
```
⚠️ Note: `container/list` is idiomatic for linked lists in Go.  
Your Piscine solution demonstrates how to implement a custom singly linked list manually.

## Skills Practiced
- Structs and pointers
- Linked list insertion at the front
- Handling empty list cases
- Updating head and tail references

## Notes
- This exercise demonstrates how to implement dynamic data structures in Go.
- For production code, prefer `container/list` unless you need a custom implementation.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Effective Go — Data structures (go.dev in Bing)