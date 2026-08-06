package main

import (
    "fmt"
)

func main() {
    root := &TreeNode{Data: "4"}
    BTreeInsertData(root, "1")
    BTreeInsertData(root, "7")
    BTreeInsertData(root, "5")

    max := BTreeMax(root)
    fmt.Println(max.Data) // 7
}

func BTreeMax(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    current := root
    for current.Right != nil {
        current = current.Right
    }
    return current
}

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
