package huffman

//Node ...
type Node struct {
	Charachter string
	Weight     uint64
	CodePath   *BitArray
	Nodes      []*Node
}

//UpdateCodePath ...
func (n *Node) UpdateCodePath(side uint) {
	n.CodePath.Prepend(side)

	for _, node := range n.Nodes {
		node.UpdateCodePath(side)
	}
}
