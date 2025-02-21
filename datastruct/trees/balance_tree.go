package trees

type T int

type Node struct {
	Value  T
	Height int
	Left   *Node
	Right  *Node
}

type BinaryBalanceTree struct {
	root *Node
}

func NewBalanceTree(value T) *BinaryBalanceTree {
	return &BinaryBalanceTree{}
}

func (n *Node) GetHeight() int {
	if n == nil {
		return 0
	}
	return n.Height
}

func (n *Node) UpdateHeight() {
	n.Height = 1 + max(n.Left.GetHeight(), n.Right.GetHeight())
}

func (n *Node) GetBalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.Left.GetHeight() - n.Right.GetHeight()
}

// 右旋
func RightRotate(n *Node) *Node {
	x := n.Left
	x2 := x.Right
	x.Right = n
	n.Left = x2
	n.UpdateHeight()
	x.UpdateHeight()
	return x
}

// 左旋
func LeftRotate(n *Node) *Node {
	x := n.Right
	x2 := x.Left
	x.Left = n
	n.Right = x2
	n.UpdateHeight()
	x.UpdateHeight()
	return x
}

func (bbt *BinaryBalanceTree) Insert(value T) {
	bbt.root = bbt.insert(bbt.root, value)
}

func (bbt *BinaryBalanceTree) insert(node *Node, value T) *Node {
	if node == nil {
		return &Node{Value: value, Height: 1}
	}

	if value < node.Value {
		node.Left = bbt.insert(node.Left, value)
	} else if value > node.Value {
		node.Right = bbt.insert(node.Right, value)
	} else {
		return node
	}

	node.UpdateHeight()

	return node
}
