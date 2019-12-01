package main

import (
	"fmt"
	"log"
)

// 注意这个全局变量
var RootNodeIndex int32

type TreeNode struct {
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
	Data   int32
}

func TreePrint(data int32) {
	fmt.Println(data)
}

func InOrderWalk(t *TreeNode, f func(int32)) {
	if t != nil {
		InOrderWalk(t.Left, f)
		f(t.Data)
		InOrderWalk(t.Right, f)
	}
}

func ReConstruct(preOrder, inOrder []int32, startIndex, endIndex int32) *TreeNode {
	log.Printf("start: %d, end: %d", startIndex, endIndex)
	if startIndex > endIndex {
		return nil
	}

	t := NewTreeNode(preOrder[RootNodeIndex])
	// 每次递归构建节点, 才能+1
	RootNodeIndex += 1
	log.Printf("rootindex: %d", RootNodeIndex)
	if startIndex == endIndex {
		return t
	}
	inIndex := searchRootIndex(inOrder, startIndex, endIndex, t.Data)
	t.Left = ReConstruct(preOrder, inOrder, startIndex, inIndex-1)
	t.Right = ReConstruct(preOrder, inOrder, inIndex+1, endIndex)
	return t
}

func searchRootIndex(inOrder []int32, startIndex, endIndex, data int32) int32 {
	for i := startIndex; i <= endIndex; i++ {
		if data == inOrder[i] {
			return i
		}
	}
	return -1
}

func NewTreeNode(data int32) *TreeNode {
	t := new(TreeNode)
	t.Data = data
	t.Left = nil
	t.Right = nil
	return t
}

func MinNode(t *TreeNode) *TreeNode {
	for {
		if t.Left != nil {
			t = t.Left
		} else {
			break
		}
	}
	return t
}

func MaxNode(t *TreeNode) *TreeNode {
	for {
		if t.Right != nil {
			t = t.Right
		} else {
			break
		}
	}
	return t
}
func Insert(t *TreeNode, data int32) *TreeNode {
	newNode := NewTreeNode(data)
	if t != nil {
		if t.Data > data {
			if t.Left != nil {
				t = t.Left
				Insert(t, data)
			} else {
				t.Left = newNode
			}
		} else if t.Data < data {
			if t.Right != nil {
				t = t.Right
				Insert(t, data)
			} else {
				t.Right = newNode
			}
		}
	} else {
		t = newNode
	}
	return t
}

func main() {
	RootNodeIndex = 0
	// preOrderArray := []int32{1, 2, 4, 3, 5, 6}
	// inOrderArray := []int32{4, 2, 1, 5, 3, 6}
	// root := ReConstruct(preOrderArray, inOrderArray, 0, int32(len(inOrderArray)-1))
	// InOrderWalk(root, TreePrint)
	bstreeInOrder := []int32{1, 2, 3, 4, 5, 6}
	bstreePreOrder := []int32{3, 2, 1, 5, 4, 6}
	bstree := ReConstruct(bstreePreOrder, bstreeInOrder, 0, int32(len(bstreeInOrder)-1))
	InOrderWalk(bstree, TreePrint)
	minNode := MinNode(bstree)
	fmt.Println(minNode.Data)

	// insert data to a tree
	t := Insert(nil, 2)
	t = Insert(t, 3)
	t = Insert(t, 1)
	t = Insert(t, 5)
	t = Insert(t, 4)
	t = Insert(t, 6)
	fmt.Println(MaxNode(t).Data)
}
