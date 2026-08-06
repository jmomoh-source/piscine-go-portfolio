# Quest11 — listmerge

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **linked list concatenation** in Go.  
The task: write a function `ListMerge` that appends all elements of list `l2` to the end of list `l1`.

Rules:
- Use the provided `NodeL` and `List` structures.
- Do not create new nodes — reuse existing ones.
- Update `Tail` of `l1` to point to the last node of `l2`.

## Instructions
- File to submit: `listmerge.go`
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

func ListMerge(l1 *List, l2 *List)
```

## Implementation
`listmerge.go`:
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

func ListMerge(l1 *List, l2 *List) {
    if l1.Head == nil {
        l1.Head = l2.Head
        l1.Tail = l2.Tail
    } else if l2.Head != nil {
        l1.Tail.Next = l2.Head
        l1.Tail = l2.Tail
    }
}
```

### Explanation
- If `l1` is empty, set its head and tail to `l2`’s head and tail.
- If `l1` is non‑empty and `l2` has elements:
  - Link `l1.Tail.Next` to `l2.Head`.
  - Update `l1.Tail` to `l2.Tail`.
- No new nodes are created; existing ones are reused.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func PrintList(l *piscine.List) {
    it := l.Head
    for it != nil {
        fmt.Print(it.Data, " -> ")
        it = it.Next
    }
    fmt.Print(nil, "\n")
}

func main() {
    link := &piscine.List{}
    link2 := &piscine.List{}

    piscine.ListPushBack(link, "a")
    piscine.ListPushBack(link, "b")
    piscine.ListPushBack(link, "c")
    piscine.ListPushBack(link, "d")
    fmt.Println("-----first List------")
    PrintList(link)

    piscine.ListPushBack(link2, "e")
    piscine.ListPushBack(link2, "f")
    piscine.ListPushBack(link2, "g")
    piscine.ListPushBack(link2, "h")
    fmt.Println("-----second List------")
    PrintList(link2)

    fmt.Println("-----Merged List-----")
    piscine.ListMerge(link, link2)
    PrintList(link)
}
```

Output:
```text
-----first List------
a -> b -> c -> d -> <nil>
-----second List------
e -> f -> g -> h -> <nil>
-----Merged List-----
a -> b -> c -> d -> e -> f -> g -> h -> <nil>
```

## Standard Library Equivalent
Go’s `container/list` package provides a doubly linked list. You can merge by iterating over one list and appending to another:
```go
import (
    "container/list"
    "fmt"
)

func main() {
    l1 := list.New()
    l1.PushBack("a")
    l1.PushBack("b")

    l2 := list.New()
    l2.PushBack("c")
    l2.PushBack("d")

    for e := l2.Front(); e != nil; e = e.Next() {
        l1.PushBack(e.Value)
    }

    for e := l1.Front(); e != nil; e = e.Next() {
        fmt.Print(e.Value, " -> ")
    }
    fmt.Println("nil")
}
```
⚠️ Note: `container/list` does not have a built‑in merge method.  
Your Piscine solution demonstrates how to implement efficient concatenation manually.

## Skills Practiced
- Structs and pointers
- Linked list concatenation
- Efficient tail management
- Reusing existing nodes

## Notes
- This exercise demonstrates how to merge two lists without creating new nodes.
- It highlights the importance of maintaining correct head and tail references.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Effective Go — Data structures (go.dev in Bing)