# Quest11 — listsize

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **linked list traversal and counting** in Go.  
The task: write a function `ListSize` that returns the number of elements in a linked list.

Rules:
- Use the provided `NodeL` and `List` structures.
- Traverse the list starting from `Head`.
- Count each node until the end (`Next == nil`).
- Return the total count.

## Instructions
- File to submit: `listsize.go`
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

func ListSize(l *List) int
```

## Implementation
`listsize.go`:
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

func ListSize(l *List) int {
    count := 0
    it := l.Head
    for it != nil {
        count++
        it = it.Next
    }
    return count
}
```

### Explanation
- Initialize a counter to `0`.
- Start from the `Head` node.
- Traverse the list using `Next`.
- Increment the counter for each node.
- Return the counter after traversal.

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
    piscine.ListPushFront(link, "2")
    piscine.ListPushFront(link, "you")
    piscine.ListPushFront(link, "man")

    fmt.Println(piscine.ListSize(link))
}
```

Output:
```text
4
```

## Standard Library Equivalent
Go’s `container/list` package provides a doubly linked list with a built‑in `Len()` method:
```go
import (
    "container/list"
    "fmt"
)

func main() {
    l := list.New()
    l.PushFront("Hello")
    l.PushFront("2")
    l.PushFront("you")
    l.PushFront("man")

    fmt.Println(l.Len()) // 4
}
```
⚠️ Note: `container/list.Len()` is concise and idiomatic.  
Your Piscine solution demonstrates how to manually traverse and count nodes in a custom singly linked list.

## Skills Practiced
- Structs and pointers
- Linked list traversal
- Counting elements
- Iteration with `for` loops

## Notes
- This exercise demonstrates how to implement basic traversal logic in linked lists.
- For production code, prefer `container/list` unless you need a custom implementation.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Effective Go — Data structures (go.dev in Bing)