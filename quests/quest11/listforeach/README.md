# Quest11 — listforeach

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **functional iteration over linked lists** in Go.  
The task: write a function `ListForEach` that applies a given function to each node in the list.

Rules:
- Use the provided `NodeL` and `List` structures.
- The function argument must accept a pointer to a `NodeL`.
- Apply the function to every node in the list.
- Include the helper functions `Add2_node` and `Subtract3_node` in the same file.

## Instructions
- File to submit: `listforeach.go`
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

func ListForEach(l *List, f func(*NodeL))

func Add2_node(node *NodeL)
func Subtract3_node(node *NodeL)
```

## Implementation
`listforeach.go`:
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

func ListForEach(l *List, f func(*NodeL)) {
    current := l.Head
    for current != nil {
        f(current)
        current = current.Next
    }
}

func Add2_node(node *NodeL) {
    switch node.Data.(type) {
    case int:
        node.Data = node.Data.(int) + 2
    case string:
        node.Data = node.Data.(string) + "2"
    }
}

func Subtract3_node(node *NodeL) {
    switch node.Data.(type) {
    case int:
        node.Data = node.Data.(int) - 3
    case string:
        node.Data = node.Data.(string) + "-3"
    }
}
```

### Explanation
- `ListForEach` traverses the list from `Head` to `Tail`.
- For each node, it applies the function `f`.
- `Add2_node` modifies node data:
  - Adds 2 if integer.
  - Appends `"2"` if string.
- `Subtract3_node` modifies node data:
  - Subtracts 3 if integer.
  - Appends `"-3"` if string.

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

    piscine.ListPushBack(link, "1")
    piscine.ListPushBack(link, "2")
    piscine.ListPushBack(link, "3")
    piscine.ListPushBack(link, "5")

    piscine.ListForEach(link, piscine.Add2_node)

    it := link.Head
    for it != nil {
        fmt.Println(it.Data)
        it = it.Next
    }
}
```

Output:
```text
12
22
32
52
```

## Standard Library Equivalent
Go’s `container/list` package allows iteration with a loop:
```go
import (
    "container/list"
    "fmt"
)

func Add2(e *list.Element) {
    switch e.Value.(type) {
    case int:
        e.Value = e.Value.(int) + 2
    case string:
        e.Value = e.Value.(string) + "2"
    }
}

func main() {
    l := list.New()
    l.PushBack("1")
    l.PushBack("2")
    l.PushBack("3")
    l.PushBack("5")

    for e := l.Front(); e != nil; e = e.Next() {
        Add2(e)
    }

    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
}
```
⚠️ Note: `container/list` is idiomatic for linked lists.  
Your Piscine solution demonstrates how to implement functional iteration manually.

## Skills Practiced
- Structs and pointers
- Linked list traversal
- Higher‑order functions
- Data mutation via callbacks

## Notes
- This exercise demonstrates how to apply functions to each node in a linked list.
- It highlights Go’s support for first‑class functions and functional programming patterns.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Effective Go — Functions (go.dev in Bing)