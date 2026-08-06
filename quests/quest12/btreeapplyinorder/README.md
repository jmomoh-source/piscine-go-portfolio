# Quest12 — btreeapplyinorder

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **in‑order traversal** of binary trees in Go.  
The task: write a function `BTreeApplyInorder` that applies a given function `f` to each node’s data in ascending order.

Rules:
- Use the `TreeNode` structure defined in `btreeinsertdata`.
- Traverse the tree in **in‑order**: Left → Root → Right.
- Apply the function `f` to each node’s `Data`.

## Instructions
- File to submit: `btreeapplyinorder.go`
- Expected function signature:
```go
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error))
```

## Implementation
`btreeapplyinorder.go`:
```go
package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
    if root == nil {
        return
    }
    BTreeApplyInorder(root.Left, f)
    f(root.Data)
    BTreeApplyInorder(root.Right, f)
}
```

### Explanation
- If the current node is `nil`, return (base case).
- Recursively traverse the left subtree.
- Apply the function `f` to the current node’s `Data`.
- Recursively traverse the right subtree.
- This ensures values are processed in ascending order for a BST.

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

    piscine.BTreeApplyInorder(root, fmt.Println)
}
```

Output:
```text
1
4
5
7
```

## Standard Library Equivalent
Go’s standard library does not provide binary tree traversal functions.  
Instead, developers typically implement recursive traversals like this one.  
Your Piscine solution demonstrates how to apply functions directly during traversal.

## Skills Practiced
- Recursive algorithms
- Tree traversal (in‑order)
- Applying functions to data
- Working with variadic function signatures

## Notes
- In‑order traversal is fundamental for BSTs because it visits nodes in ascending order.
- Later exercises will explore pre‑order, post‑order, and level‑order traversals.