package huffman

import (
	"errors"
	"sort"
)

var leftPathValue = 0
var rightPathValue = 1

//Compress ...
func Compress(text string) ([]byte, *Node) {
	characters := make(map[string]int64)

	//go through text and get characters and frwquencies
	for _, c := range text {
		char := string(c)
		characters[char]++
	}

	//create tree
	huffmanTree := createHuffmanTree(characters)

	//construct codes from tree
	compressed := make([]byte, 0)
	for _, c := range text {
		char := string(c)
		code, err := GetCodePathForChar(huffmanTree, char, make([]byte, 0))
		if err != nil {
			panic(err)
		}
		compressed = append(compressed[:], code[:]...)
	}

	return compressed, huffmanTree
}

//Decompress ...
func Decompress(compressed []byte, huffmanTree *Node) string {

	var pointer int64
	var decompressed = ""
	pointer = 0
	for pointer < int64(len(compressed)) {
		var char string

		//find first matching character and move pointer forward
		char, pointer = FindCharByCode(compressed, pointer, huffmanTree)
		decompressed += char
		pointer++
	}

	return decompressed
}

func createHuffmanTree(characters map[string]int64) *Node {
	nodes := make([]*Node, 0)
	for key, value := range characters {
		node := &Node{}
		node.Charachter = key
		node.Weight = value

		nodes = append(nodes, node)
	}

	sortNodesByWeight(&nodes)

	for len(nodes) > 1 {
		newNode := &Node{}

		var first, second Node
		first, nodes = *nodes[0], nodes[1:]
		second, nodes = *nodes[0], nodes[1:]

		newNode.Left = &first
		newNode.Right = &second

		//sum weights
		newNode.Weight = first.Weight + second.Weight

		nodes = append(nodes, newNode)
		sortNodesByWeight(&nodes)
	}

	return nodes[0]
}

func sortNodesByWeight(nodes *[]*Node) {

	//sort first by name
	sort.Slice(*nodes, func(i, j int) bool {
		return (*nodes)[i].Charachter < (*nodes)[j].Charachter
	})

	//sort by weight
	sort.Slice(*nodes, func(i, j int) bool {
		return (*nodes)[i].Weight < (*nodes)[j].Weight
	})
}

//GetCodePathForChar ...
func GetCodePathForChar(n *Node, char string, code []byte) ([]byte, error) {
	if n == nil {
		return code[0 : len(code)-1], errors.New("Reached to end leaf without matching character")
	}

	if n.Left != nil && n.Left.Charachter == char {
		code = append(code, 0)
	} else if n.Right != nil && n.Right.Charachter == char {
		code = append(code, 1)
	} else {
		var err error
		code = append(code, 0)

		code, err = GetCodePathForChar(n.Left, char, code)

		if err != nil {
			code = append(code, 1)
			code, err = GetCodePathForChar(n.Right, char, code)
			if err != nil {
				return code[0 : len(code)-1], err
			}
		}
	}

	return code, nil
}

//FindCharByCode ...
func FindCharByCode(codes []byte, pointer int64, n *Node) (string, int64) {
	code := codes[pointer]

	if code == 0 {
		if n.Left != nil && len(n.Left.Charachter) > 0 {
			return n.Left.Charachter, pointer
		}
		pointer++
		return FindCharByCode(codes, pointer, n.Left)
	} else if code == 1 {
		if n.Right != nil && len(n.Right.Charachter) > 0 {
			return n.Right.Charachter, pointer
		}

		pointer++
		return FindCharByCode(codes, pointer, n.Right)
	}
	panic("Code " + string(code) + " is not supported")
}

//Node ...
type Node struct {
	Charachter string
	Weight     int64
	Left       *Node
	Right      *Node
}
