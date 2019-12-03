package binary_tree

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

// this func is buggy
func Insert(t *TreeNode, data int32) *TreeNode {
	newNode := NewTreeNode(data)
	if t != nil {
		if t.Data > data {
			if t.Left != nil {
				Insert(t.Left, data)
			} else {
				t.Left = newNode
			}
		} else if t.Data < data {
			if t.Right != nil {
				Insert(t.Right, data)
			} else {
				t.Right = newNode
			}
		}
	} else {
		t = newNode
	}
	return t
}

func SliceToBSTree(sliceData []int32) *TreeNode {
	/*
		// 先声明t的类型, 是一个指针
		t := new(TreeNode)
		// 再赋值空指针
		t = nil
	*/
	// 或者
	var t *TreeNode
	for i := 0; i < len(sliceData); i++ {
		t = Insert(t, sliceData[i])
	}
	return t
}

func BSTreeToSlice(t *TreeNode, s []int32) []int32 {
	if t != nil {
		s = BSTreeToSlice(t.Left, s)
		s = append(s, t.Data)
		s = BSTreeToSlice(t.Right, s)
	}
	return s
}

func LookUp(t *TreeNode, data int32) *TreeNode {
	for {
		if t != nil && t.Data != data {
			if data < t.Data {
				t = t.Left
			} else {
				t = t.Right
			}
		} else {
			break
		}
	}
	return t
}

func Succ(t *TreeNode) *TreeNode {
	if t.Right != nil {
		return MinNode(t.Right)
	} else {
		p := t.Parent
		for {
			if p != nil && p.Left != t { //直到找到出现左子树
				t = p
				p = p.Parent
			} else {
				break
			}
		}
		return p
	}
}

func Pred(t *TreeNode) *TreeNode {
	if t.Left != nil {
		return MaxNode(t.Left)
	} else {
		p := t.Parent
		for {
			if p != nil && p.Right != t { //直到找到出现右子树
				t = p
				p = p.Parent
			} else {
				break
			}
		}
		return p
	}
}

// 如果二叉树是左斜的, 搜索根节点和递归累计, 时间复杂度为O(n^2)
// 改进: 搜索根节点可以改成用map存储时间复杂度提升到O(1), 总时间复杂度提升到O(n)
