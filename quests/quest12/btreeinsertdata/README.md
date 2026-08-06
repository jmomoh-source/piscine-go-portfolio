# Quest12 — btreeinsertdata

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **insertion in binary search trees (BSTs)** in Go.  
The task: write a function `BTreeInsertData` that inserts new data into a BST while maintaining its ordering properties.

Rules:
- Use the provided `TreeNode` structure.
- Insert nodes such that:
  - Values smaller than the current node go to the left.
  - Values greater than or equal go to the right.
- Maintain parent pointers for each inserted node.

## Instructions
- File to submit: `btreeinsertdata.go`
- Expected structure and function signature:
```go
type TreeNode struct {
    Left, Right, Parent *TreeNode
    Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode
```

## Implementation
`btreeinsertdata.go`:
```go
package piscine

type TreeNode struct {
    Left, Right, Parent *TreeNode
    Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
    if root == nil {
        return &TreeNode{Data: data}
    }

    if data < root.Data {
        if root.Left == nil {
            root.Left = &TreeNode{Data: data, Parent: root}
        } else {
            BTreeInsertData(root.Left, data)
        }
    } else {
        if root.Right == nil {
            root.Right = &TreeNode{Data: data, Parent: root}
        } else {
            BTreeInsertData(root.Right, data)
        }
    }

    return root
}
```

### Explanation
- If the tree is empty, create a new root node.
- Compare `data` with the current node’s `Data`:
  - If smaller, insert into the left subtree.
  - If greater or equal, insert into the right subtree.
- Maintain the `Parent` pointer when creating new nodes.
- Return the root of the tree.

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

    fmt.Println(root.Left.Data)       // 1
    fmt.Println(root.Data)            // 4
    fmt.Println(root.Right.Left.Data) // 5
    fmt.Println(root.Right.Data)      // 7
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
Go’s standard library does not provide a binary tree implementation.  
Instead, developers typically use slices with `sort` or implement custom tree structures.  
Your Piscine solution demonstrates how to build and manipulate a BST directly.

## Skills Practiced
- Structs and pointers
- Binary search tree insertion
- Maintaining parent references
- Recursive algorithms

## Notes
- This exercise sets the foundation for subsequent tree operations (search, traversals, deletion).
- Parent pointers are useful for advanced operations like transplant and delete.