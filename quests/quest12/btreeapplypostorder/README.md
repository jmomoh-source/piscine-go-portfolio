# Quest12 — btreeapplypostorder

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **post‑order traversal** of binary trees in Go.  
The task: write a function `BTreeApplyPostorder` that applies a given function `f` to each node’s data using a post‑order walk.

Rules:
- Use the `TreeNode` structure defined in `btreeinsertdata`.
- Traverse the tree in **post‑order**: Left → Right → Root.
- Apply the function `f` to each node’s `Data`.

## Instructions
- File to submit: `btreeapplypostorder.go`
- Expected function signature:
```go
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error))
```

## Implementation
`btreeapplypostorder.go`:
```go
package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
    if root == nil {
        return
    }
    BTreeApplyPostorder(root.Left, f)
    BTreeApplyPostorder(root.Right, f)
    f(root.Data)
}
```

### Explanation
- If the current node is `nil`, return (base case).
- Recursively traverse the left subtree.
- Recursively traverse the right subtree.
- Apply the function `f` to the current node’s `Data`.
- This ensures values are processed in **post‑order**.

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

    piscine.BTreeApplyPostorder(root, fmt.Println)
}
```

Output:
```text
1
5
7
4
```

## Standard Library Equivalent
Go’s standard library does not provide binary tree traversal functions.  
Instead, developers typically implement recursive traversals like this one.  
Your Piscine solution demonstrates how to apply functions directly during traversal.

## Skills Practiced
- Recursive algorithms
- Tree traversal (post‑order)
- Applying functions to data
- Working with variadic function signatures

## Notes
- Post‑order traversal is often used for deleting or freeing nodes because it processes children before the parent.
- Later exercises will explore pre‑order and level‑order traversals.