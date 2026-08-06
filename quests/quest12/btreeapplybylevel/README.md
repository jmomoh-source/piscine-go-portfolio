# Quest12 — btreeapplybylevel

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **level‑order traversal (breadth‑first search)** of binary trees in Go.  
The task: write a function `BTreeApplyByLevel` that applies a given function `f` to each node’s data, level by level.

Rules:
- Use the `TreeNode` structure defined in `btreeinsertdata`.
- Traverse the tree level by level (root first, then children).
- Apply the function `f` to each node’s `Data`.

## Instructions
- File to submit: `btreeapplybylevel.go`
- Expected function signature:
```go
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error))
```

## Implementation
`btreeapplybylevel.go`:
```go
package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
    if root == nil {
        return
    }

    queue := []*TreeNode{root}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        f(node.Data)

        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
}
```

### Explanation
- Use a slice as a queue to store nodes level by level.
- Start with the root node.
- For each node:
  - Apply the function `f` to its `Data`.
  - Enqueue its left and right children if they exist.
- Continue until the queue is empty.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    root := &piscine.TreeNode{Data: "4"}
    piscine.BTreeInsertData(root, "1")
    piscine.BTreeInsertData(root, "7")
    piscine.BTreeInsertData(root, "5")

    piscine.BTreeApplyByLevel(root, fmt.Println)
}
```

Output:
```text
4
1
7
5
```

## Standard Library Equivalent
Go’s standard library does not provide a binary tree type or level‑order traversal.  
This queue‑based approach is the standard way to implement breadth‑first traversal.

## Skills Practiced
- Breadth‑first search (BFS)
- Queue data structure
- Tree traversal
- Applying functions to hierarchical data

## Notes
- Level‑order traversal is useful for algorithms like printing trees by depth or checking completeness.
- This exercise complements in‑order, pre‑order, and post‑order traversals by covering breadth‑first traversal.