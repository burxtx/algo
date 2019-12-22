package binary_tree

type RBNode struct {
	Parent *RBNode
	Left   *RBNode
	Right  *RBNode
	Data   int32
	Color  int32
}

const (
	RED   = 1
	BLACK = 2
)

func (n *RBNode) SetLeft(x *RBNode) {
	n.Left = x
	if x != nil {
		x.Parent = n
	}
}

func (n *RBNode) SetRight(x *RBNode) {
	n.Right = x
	if x != nil {
		x.Parent = n
	}
}

func (n *RBNode) SetChildren(x, y *RBNode) {
	n.SetLeft(x)
	n.SetRight(y)
}
func (n *RBNode) Sibling() *RBNode {
	if n.Parent.Left == n {
		return n.Parent.Right
	} else {
		return n.Parent.Left
	}
}
func (n *RBNode) Uncle() *RBNode {
	return n.Parent.Sibling()
}

func (n *RBNode) Grandparent() *RBNode {
	return n.Parent.Parent
}

func LeftRotate(t, x *RBNode) {

}

func SetColor(nodes []*RBNode, colors []int32) {
	if len(nodes) != len(colors) {
		return
	}
	for i := 0; i < len(nodes); i++ {
		nodes[i].Color = colors[i]
	}
}

func RBInsert(t *RBNode, data int32) *RBNode {
	root := t
	x := &RBNode{Data: data, Color: RED}
	var parent *RBNode
	// 找到插入位置，插入位置总是叶子节点
	for {
		if t != nil {
			parent = t
			if data < t.Data {
				t = t.Left
			} else {
				t = t.Right
			}
		} else {
			break
		}
	}
	if parent == nil {
		root = x
	} else if data < parent.Data {
		parent.SetLeft(x)
	} else {
		parent.SetRight(x)
	}
	return RBInsertFix(root, x)
}

func RBInsertFix(t, x *RBNode) *RBNode {

}
