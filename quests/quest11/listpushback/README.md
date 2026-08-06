# Quest11 — listpushback

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **linked list insertion at the back** in Go.  
The task: write a function `ListPushBack` that inserts a new node at the end of a linked list.

Rules:
- Use the provided `NodeL` and `List` structures.
- Update both `Head` and `Tail` pointers correctly.
- Handle the case where the list is initially empty.

## Instructions
- File to submit: `listpushback.go`
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

func ListPushBack(l *List, data interface{})
```

## Implementation
`listpushback.go`:
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

func ListPushBack(l *List, data interface{}) {
    newNode := &NodeL{Data: data}
    if l.Head == nil {
        l.Head = newNode
        l.Tail = newNode
    } else {
        l.Tail.Next = newNode
        l.Tail = newNode
    }
}
```

### Explanation
- Create a new node with the given data.
- If the list is empty (`Head == nil`), set both `Head` and `Tail` to the new node.
- Otherwise, link the current `Tail` to the new node and update `Tail`.

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

    piscine.ListPushBack(link, "Hello")
    piscine.ListPushBack(link, "man")
    piscine.ListPushBack(link, "how are you")

    for link.Head != nil {
        fmt.Println(link.Head.Data)
        link.Head = link.Head.Next
    }
}
```

Output:
```text
Hello
man
how are you
```

## Standard Library Equivalent
Go does not have a built‑in linked list type, but the `container/list` package provides a doubly linked list:
```go
import (
    "container/list"
    "fmt"
)

func main() {
    l := list.New()
    l.PushBack("Hello")
    l.PushBack("man")
    l.PushBack("how are you")

    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
}
```
⚠️ Note: `container/list` is idiomatic for linked lists in Go.  
Your Piscine solution demonstrates how to implement a custom singly linked list manually.

## Skills Practiced
- Structs and pointers
- Linked list insertion
- Handling empty list cases
- Updating head and tail references

## Notes
- This exercise demonstrates how to implement dynamic data structures in Go.
- For production code, prefer `container/list` unless you need a custom implementation.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Effective Go — Data structures (go.dev in Bing) 