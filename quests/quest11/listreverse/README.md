# Quest11 — listreverse

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **linked list reversal** in Go.  
The task: write a function `ListReverse` that reverses the order of the elements in a linked list.

Rules:
- Use the provided `NodeL` and `List` structures.
- Reverse the list in place by re‑linking nodes.
- Update both `Head` and `Tail` pointers correctly.

## Instructions
- File to submit: `listreverse.go`
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

func ListReverse(l *List)
```

## Implementation
`listreverse.go`:
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

func ListReverse(l *List) {
    var prev *NodeL
    current := l.Head
    l.Tail = l.Head
    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }
    l.Head = prev
}
```

### Explanation
- Initialize `prev` as `nil` and `current` as the head.
- Iterate through the list:
  - Save the next node.
  - Reverse the link (`current.Next = prev`).
  - Move `prev` and `current` forward.
- After traversal, set `Head` to `prev` (new head).
- Update `Tail` to the old head.

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

    piscine.ListPushBack(link, 1)
    piscine.ListPushBack(link, 2)
    piscine.ListPushBack(link, 3)
    piscine.ListPushBack(link, 4)

    piscine.ListReverse(link)

    it := link.Head
    for it != nil {
        fmt.Println(it.Data)
        it = it.Next
    }

    fmt.Println("Tail", link.Tail)
    fmt.Println("Head", link.Head)
}
```

Output:
```text
4
3
2
1
Tail &{1 <nil>}
Head &{4 0xc42000a140}
```

## Standard Library Equivalent
Go’s `container/list` package is doubly linked, so you can traverse backwards using `Back()`:
```go
import (
    "container/list"
    "fmt"
)

func main() {
    l := list.New()
    l.PushBack(1)
    l.PushBack(2)
    l.PushBack(3)
    l.PushBack(4)

    for e := l.Back(); e != nil; e = e.Prev() {
        fmt.Println(e.Value)
    }
}
```
⚠️ Note: `container/list` allows reverse traversal but does not provide an in‑place reversal method.  
Your Piscine solution demonstrates how to implement reversal manually in a singly linked list.

## Skills Practiced
- Structs and pointers
- Linked list traversal
- In‑place reversal algorithms
- Updating head and tail references

## Notes
- This exercise demonstrates how to reverse a linked list efficiently.
- Unlike arrays or slices, reversal in linked lists requires re‑linking nodes.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Effective Go — Data structures (go.dev in Bing)