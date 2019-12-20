package binary_tree

import (
	"fmt"
	"strings"
	"testing"
)

func TestRecurseInsert(t *testing.T) {
	RootNodeIndex = 0
	preOrderArray := []int32{1, 2, 4, 3, 5, 6}
	inOrderArray := []int32{4, 2, 1, 5, 3, 6}
	tree := ReConstruct(preOrderArray, inOrderArray, 0, int32(len(inOrderArray)-1))
	// insert a node
	var x int32
	x = 7
	newT := Insert(tree, x)
	InOrderWalk(newT, TreePrint)
}

func TestInsert(t *testing.T) {
	// insert data to a tree
	tree := Insert(nil, 2)
	tree = Insert(tree, 3)
	tree = Insert(tree, 1)
	tree = Insert(tree, 5)
	tree = Insert(tree, 4)
	tree = Insert(tree, 6)
	maxNode := MaxNode(tree).Data
	if maxNode != 6 {
		t.Errorf("max node is %d, wanted %d", maxNode, 6)
	}
}

func SliceToString(a []int32, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
func TestSliceToTree(t *testing.T) {
	tests := []struct {
		input  []int32
		wanted []int32
	}{
		{[]int32{2, 3, 1, 5, 4, 6}, []int32{1, 2, 3, 4, 5, 6}},
	}
	for _, test := range tests {
		tree := SliceToBSTree(test.input)
		s := make([]int32, 0)
		s = BSTreeToSlice(tree, s)
		fmt.Println(s)
		fmt.Println(SliceToString(s, ""))
		if SliceToString(s, "") != SliceToString(test.wanted, "") {
			t.Error("actual    wanted")
			t.Errorf("%v != %v", SliceToString(s, ""), SliceToString(test.wanted, ""))
		}
	}
}

func TestSucc(t *testing.T) {
	tests := []struct {
		inputTree []int32
		lookup    int32
		wanted    int32
	}{
		{[]int32{2, 3, 1, 5, 4, 6}, 3, 4},
	}
	for _, test := range tests {
		tree := SliceToBSTree(test.inputTree)
		res := Succ(LookUp(tree, test.lookup))
		fmt.Println(res)
		if res.Data != test.wanted {
			t.Errorf("%d!=%d", res.Data, test.wanted)
		}
	}
}
