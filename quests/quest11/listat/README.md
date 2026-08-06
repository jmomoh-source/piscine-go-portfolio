# Quest11 — listat

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **linked list indexed access** in Go.  
The task: write a function `ListAt` that returns the pointer to the node at a given position in the linked list.

Rules:
- Use the provided `NodeL` structure.
- Traverse the list starting from `Head`.
- Return the node pointer at the given position.
- If the position is invalid (negative or beyond list length), return `nil`.

## Instructions
- File to submit: `listat.go`
- Expected structure and function signature:
```go
type NodeL struct {
    Data interface{}
    Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL
```

## Implementation
`listat.go`:
```go
package piscine

type NodeL struct {
    Data interface{}
    Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
    if pos < 0 {
        return nil
    }
    index := 0
    current := l
    for current != nil {
        if index == pos {
            return current
        }
        current = current.Next
        index++
    }
    return nil
}
```

### Explanation
- If `pos` is negative, return `nil`.
- Start from the head node.
- Traverse the list, incrementing an index counter.
- When the counter equals `pos`, return the current node.
- If traversal ends before reaching `pos`, return `nil`.

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
    piscine.ListPushBack(link, "how are")
    piscine.ListPushBack(link, "you")
    piscine.ListPushBack(link, 1)

    fmt.Println(piscine.ListAt(link.Head, 3).Data)
    fmt.Println(piscine.ListAt(link.Head, 1).Data)
    fmt.Println(piscine.ListAt(link.Head, 7))
}
```

Output:
```text
1
how are
<nil>
```

## Standard Library Equivalent
Go’s `container/list` package does not provide direct indexed access, but you can simulate it by traversing manually:
```go
import (
    "container/list"
    "fmt"
)

func ElementAt(l *list.List, pos int) *list.Element {
    if pos < 0 {
        return nil
    }
    index := 0
    for e := l.Front(); e != nil; e = e.Next() {
        if index == pos {
            return e
        }
        index++
    }
    return nil
}

func main() {
    l := list.New()
    l.PushBack("hello")
    l.PushBack("how are")
    l.PushBack("you")
    l.PushBack(1)

    fmt.Println(ElementAt(l, 3).Value) // 1
    fmt.Println(ElementAt(l, 1).Value) // how are
    fmt.Println(ElementAt(l, 7))       // <nil>
}
```
⚠️ Note: `container/list` is idiomatic for linked lists, but indexed access requires manual traversal.  
Your Piscine solution demonstrates how to implement this logic explicitly.

## Skills Practiced
- Structs and pointers
- Linked list traversal
- Indexed access
- Error handling for invalid positions

## Notes
- This exercise demonstrates how to retrieve nodes by index in a linked list.
- Unlike slices, linked lists do not support O(1) indexed access; traversal is required.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Effective Go — Data structures (go.dev in Bing) 