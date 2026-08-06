# Quest12 — btreeapplypreorder

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **pre‑order traversal** of binary trees in Go.  
The task: write a function `BTreeApplyPreorder` that applies a given function `f` to each node’s data using a pre‑order walk.

Rules:
- Use the `TreeNode` structure defined in `btreeinsertdata`.
- Traverse the tree in **pre‑order**: Root → Left → Right.
- Apply the function `f` to each node’s `Data`.

## Instructions
- File to submit: `btreeapplypreorder.go`
- Expected function signature:
```go
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error))
```

## Implementation
`btreeapplypreorder.go`:
```go
package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
    if root == nil {
        return
    }
    f(root.Data)
    BTreeApplyPreorder(root.Left, f)
    BTreeApplyPreorder(root.Right, f)
}
```

### Explanation
- If the current node is `nil`, return (base case).
- Apply the function `f` to the current node’s `Data`.
- Recursively traverse the left subtree.
- Recursively traverse the right subtree.
- This ensures values are processed in **pre‑order**.

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

    piscine.BTreeApplyPreorder(root, fmt.Println)
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
Go’s standard library does not provide binary tree traversal functions.  
Instead, developers typically implement recursive traversals like this one.  
Your Piscine solution demonstrates how to apply functions directly during traversal.

## Skills Practiced
- Recursive algorithms
- Tree traversal (pre‑order)
- Applying functions to data
- Working with variadic function signatures

## Notes
- Pre‑order traversal is often used to **copy trees** or serialize them because it processes the root before its children.
- Later exercises will explore in‑order, post‑order, and level‑order traversals.