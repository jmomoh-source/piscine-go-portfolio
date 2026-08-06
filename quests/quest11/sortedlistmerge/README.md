# Quest11 — sortedlistmerge

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **merging two sorted linked lists** in Go.  
The task: write a function `SortedListMerge` that merges two sorted lists (`n1` and `n2`) into a single sorted list.

Rules:
- Use the `NodeI` structure defined in the `listsort` exercise.
- Both input lists are already sorted.
- Merge them into one sorted list without creating new nodes.
- Return the head of the merged list.

## Instructions
- File to submit: `sortedlistmerge.go`
- Expected structure and function signature:
```go
type NodeI struct {
    Data int
    Next *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI
```

## Implementation
`sortedlistmerge.go`:
```go
package piscine

type NodeI struct {
    Data int
    Next *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
    if n1 == nil {
        return n2
    }
    if n2 == nil {
        return n1
    }

    var head *NodeI
    if n1.Data < n2.Data {
        head = n1
        n1 = n1.Next
    } else {
        head = n2
        n2 = n2.Next
    }

    current := head
    for n1 != nil && n2 != nil {
        if n1.Data < n2.Data {
            current.Next = n1
            n1 = n1.Next
        } else {
            current.Next = n2
            n2 = n2.Next
        }
        current = current.Next
    }

    if n1 != nil {
        current.Next = n1
    } else {
        current.Next = n2
    }

    return head
}
```

### Explanation
- Handle edge cases: if one list is empty, return the other.
- Choose the smaller head between `n1` and `n2` as the merged head.
- Traverse both lists, linking the smaller node each time.
- Append the remaining nodes when one list is exhausted.
- Return the merged head.

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
    var link2 *piscine.NodeI

    link = listPushBack(link, 3)
    link = listPushBack(link, 5)
    link = listPushBack(link, 7)

    link2 = listPushBack(link2, -2)
    link2 = listPushBack(link2, 9)

    PrintList(piscine.SortedListMerge(link2, link))
}
```

Output:
```text
-2 -> 3 -> 5 -> 7 -> 9 -> <nil>
```

## Standard Library Equivalent
Go’s `container/list` package does not provide a built‑in sorted merge.  
You can implement it manually:
```go
import (
    "container/list"
    "fmt"
    "sort"
)

func SortedMerge(l1, l2 *list.List) *list.List {
    var values []int
    for e := l1.Front(); e != nil; e = e.Next() {
        values = append(values, e.Value.(int))
    }
    for e := l2.Front(); e != nil; e = e.Next() {
        values = append(values, e.Value.(int))
    }
    sort.Ints(values)

    merged := list.New()
    for _, v := range values {
        merged.PushBack(v)
    }
    return merged
}

func main() {
    l1 := list.New()
    l1.PushBack(3)
    l1.PushBack(5)
    l1.PushBack(7)

    l2 := list.New()
    l2.PushBack(-2)
    l2.PushBack(9)

    merged := SortedMerge(l1, l2)
    for e := merged.Front(); e != nil; e = e.Next() {
        fmt.Print(e.Value, " -> ")
    }
    fmt.Println("nil")
}
```
⚠️ Note: `container/list` requires manual merging logic.  
Your Piscine solution demonstrates how to merge sorted singly linked lists directly.

## Skills Practiced
- Structs and pointers
- Linked list traversal
- Sorted merging
- Efficient pointer manipulation

## Notes
- This exercise demonstrates how to merge two sorted linked lists without creating new nodes.
- It highlights the importance of pointer manipulation in linked list algorithms.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Go `sort` package — Official Docs (go.dev in Bing)