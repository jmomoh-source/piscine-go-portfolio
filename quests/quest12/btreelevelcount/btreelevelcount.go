package main

import (
    "fmt"
)

func main() {
    root := &TreeNode{Data: "4"}
    BTreeInsertData(root, "1")
    BTreeInsertData(root, "7")
    BTreeInsertData(root, "5")

    fmt.Println(BTreeLevelCount(root)) // 3
}

func BTreeLevelCount(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftHeight := BTreeLevelCount(root.Left)
    rightHeight := BTreeLevelCount(root.Right)

    if leftHeight > rightHeight {
        return leftHeight + 1
    }
    return rightHeight + 1
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
