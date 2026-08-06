package main

import (
    "fmt"
)

func main() {
    root := &TreeNode{Data: "4"}
    BTreeInsertData(root, "1")
    BTreeInsertData(root, "7")
    BTreeInsertData(root, "5")

    selected := BTreeSearchItem(root, "7")
    fmt.Print("Item selected -> ")
    if selected != nil {
        fmt.Println(selected.Data)
    } else {
        fmt.Println("nil")
    }

    fmt.Print("Parent of selected item -> ")
    if selected.Parent != nil {
        fmt.Println(selected.Parent.Data)
    } else {
        fmt.Println("nil")
    }

    fmt.Print("Left child of selected item -> ")
    if selected.Left != nil {
        fmt.Println(selected.Left.Data)
    } else {
        fmt.Println("nil")
    }

    fmt.Print("Right child of selected item -> ")
    if selected.Right != nil {
        fmt.Println(selected.Right.Data)
    } else {
        fmt.Println("nil")
    }
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
    if root == nil {
        return nil
    }
    if elem == root.Data {
        return root
    } else if elem < root.Data {
        return BTreeSearchItem(root.Left, elem)
    } else {
        return BTreeSearchItem(root.Right, elem)
    }
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
