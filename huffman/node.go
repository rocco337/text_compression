package huffman

import "text_compression/bitarray"

//Node ...
type Node struct {
	Charachter string
	Weight     uint
	CodePath   *bitarray.BitArray
	Nodes      []*Node
}

//UpdateCodePath ...
func (n *Node) UpdateCodePath(side uint) {
	n.CodePath.Prepend(side)

	for _, node := range n.Nodes {
		node.UpdateCodePath(side)
	}
}
