package huffman

//Node ...
type Node struct {
	Charachter string
	Weight     int64
	CodePath   *BitArray
	Nodes      []*Node
}

//UpdateCodePath ...
func (n *Node) UpdateCodePath(side uint32) {
	n.CodePath.Preprend(side)

	for _, node := range n.Nodes {
		node.UpdateCodePath(side)
	}
}
