package huffman

import (
	"errors"
	"sort"
)

var leftPathValue = "0"
var rightPathValue = "1"

//Compress ...
func Compress(text string) (string, *Node) {
	characters := make(map[string]int64)

	//go through text and get characters and frwquencies
	for _, c := range text {
		char := string(c)
		characters[char]++
	}

	//create tree
	huffmanTree := createHuffmanTree(characters)

	//construc codes from tree
	codes := ""
	for _, c := range text {
		char := string(c)
		code, err := GetCodePathForChar(huffmanTree, char, "")
		if err != nil {
			panic(err)
		}
		codes += code
	}

	return codes, huffmanTree
}

//Decompress ...
func Decompress(compressed string, huffmanTree *Node) string {
	codes := make([]string, 0)
	for _, c := range compressed {
		char := string(c)
		codes = append(codes, char)
	}

	var i int64
	var characters = ""
	i = 0
	for i < int64(len(codes)) {
		var char string
		char, i = FindCharByCode(codes, i, huffmanTree)
		characters += char
		i++
	}

	return characters

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

	sort.Slice(*nodes, func(i, j int) bool {
		return (*nodes)[i].Charachter < (*nodes)[j].Charachter
	})

	sort.Slice(*nodes, func(i, j int) bool {
		return (*nodes)[i].Weight < (*nodes)[j].Weight
	})
}

//GetCodePathForChar ...
func GetCodePathForChar(n *Node, char string, code string) (string, error) {
	if n == nil {
		return code[0 : len(code)-1], errors.New("Reached to end leaf without matching character")
	}

	if n.Left != nil && n.Left.Charachter == char {
		code += leftPathValue
	} else if n.Right != nil && n.Right.Charachter == char {
		code += rightPathValue
	} else {
		var err error
		code, err = GetCodePathForChar(n.Left, char, code+leftPathValue)

		if err != nil {
			code, err = GetCodePathForChar(n.Right, char, code+rightPathValue)
			if err != nil {
				return code[0 : len(code)-1], err
			}
		}
	}

	return code, nil
}

//FindCharByCode ...
func FindCharByCode(codes []string, codeIndex int64, n *Node) (string, int64) {
	code := codes[codeIndex]

	if code == leftPathValue {
		if n.Left != nil && len(n.Left.Charachter) > 0 {
			return n.Left.Charachter, codeIndex
		}
		codeIndex++

		return FindCharByCode(codes, codeIndex, n.Left)
	} else if code == rightPathValue {

		if n.Right != nil && len(n.Right.Charachter) > 0 {
			return n.Right.Charachter, codeIndex
		}
		codeIndex++

		return FindCharByCode(codes, codeIndex, n.Right)
	}
	panic("Code " + code + " is not supported")
}

//Node ...
type Node struct {
	Charachter string
	Weight     int64
	Left       *Node
	Right      *Node
}
