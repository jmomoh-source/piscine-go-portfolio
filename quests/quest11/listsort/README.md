# Quest11 — listsort

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **sorting linked lists** in Go.  
The task: write a function `ListSort` that sorts the nodes of a linked list in ascending order.

Rules:
- Use the provided `NodeI` structure (with `Data int`).
- Sort the list in place by re‑ordering node values.
- Return the new head of the sorted list.

## Instructions
- File to submit: `listsort.go`
- Expected structure and function signature:
```go
type NodeI struct {
    Data int
    Next *NodeI
}

func ListSort(l *NodeI) *NodeI
```

## Implementation
`listsort.go`:
```go
package piscine

type NodeI struct {
    Data int
    Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
    if l == nil {
        return nil
    }
    for i := l; i != nil; i = i.Next {
        for j := i.Next; j != nil; j = j.Next {
            if i.Data > j.Data {
                i.Data, j.Data = j.Data, i.Data
            }
        }
    }
    return l
}
```

### Explanation
- Use a simple **bubble sort / selection sort style** algorithm:
  - Traverse the list with two pointers (`i` and `j`).
  - Compare values and swap if out of order.
- Sorting is done in place by swapping `Data` values.
- Return the head of the sorted list.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func PrintList(l *piscine.NodeI) {
    it := l
    for it != nil {
        fmt.Print(it.Data, " -> ")
        it = it.Next
    }
    fmt.Print(nil, "\n")
}

func listPushBack(l *piscine.NodeI, data int) *piscine.NodeI {
    n := &piscine.NodeI{Data: data}
    if l == nil {
        return n
    }
    iterator := l
    for iterator.Next != nil {
        iterator = iterator.Next
    }
    iterator.Next = n
    return l
}

func main() {
    var link *piscine.NodeI

    link = listPushBack(link, 5)
    link = listPushBack(link, 4)
    link = listPushBack(link, 3)
    link = listPushBack(link, 2)
    link = listPushBack(link, 1)

    PrintList(piscine.ListSort(link))
}
```

Output:
```text
1 -> 2 -> 3 -> 4 -> 5 -> <nil>
```

## Standard Library Equivalent
Go’s `container/list` package does not provide a built‑in sort method.  
Sorting requires manual traversal and re‑ordering:
```go
import (
    "container/list"
    "fmt"
    "sort"
)

func main() {
    l := list.New()
    l.PushBack(5)
    l.PushBack(4)
    l.PushBack(3)
    l.PushBack(2)
    l.PushBack(1)

    // Extract values into a slice
    var values []int
    for e := l.Front(); e != nil; e = e.Next() {
        values = append(values, e.Value.(int))
    }

    // Sort slice
    sort.Ints(values)

    // Reassign sorted values back into list
    e := l.Front()
    for _, v := range values {
        e.Value = v
        e = e.Next()
    }

    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Print(e.Value, " -> ")
    }
    fmt.Print("nil\n")
}
```
⚠️ Note: `container/list` requires manual sorting logic.  
Your Piscine solution demonstrates how to implement sorting directly on a singly linked list.

## Skills Practiced
- Structs and pointers
- Linked list traversal
- Sorting algorithms
- In‑place data manipulation

## Notes
- This exercise demonstrates how to sort linked lists without converting them to slices.
- For production code, it’s often simpler to extract values into a slice, sort with `sort.Ints`, and rebuild the list.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Go `sort` package — Official Docs (go.dev in Bing)