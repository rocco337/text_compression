package huffman

import (
	"container/heap"
	"errors"
	"fmt"
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

	//construct codes from tree
	compressed := ""
	for _, c := range text {
		char := string(c)
		code, err := GetCodePathForChar(huffmanTree, char, "")
		if err != nil {
			panic(err)
		}
		compressed += code
	}

	return compressed, huffmanTree
}

//Decompress ...
func Decompress(compressed string, huffmanTree *Node) string {

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
	nodes := make(PriorityQueue, len(characters))

	i := 0
	for _, key := range getSortedMapKeys(characters) {
		node := &Node{}
		node.Charachter = key
		node.Weight = characters[key]

		nodes[i] = &Item{
			value: node,
			index: i,
		}
		i++
	}
	heap.Init(&nodes)

	for len(nodes) > 1 {
		newNode := &Node{}

		newNode.Left = nodes.Pop().(*Node)
		newNode.Right = nodes.Pop().(*Node)

		//sum weights
		newNode.Weight = newNode.Left.Weight + newNode.Right.Weight

		newPriorityQueueItem := &Item{
			value: newNode,
		}

		heap.Push(&nodes, newPriorityQueueItem)
		nodes.update(newPriorityQueueItem, newPriorityQueueItem.value)
	}

	return nodes[0].value
}

func getSortedMapKeys(characters map[string]int64) []string {
	keys := make([]string, 0, len(characters))
	for k, _ := range characters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
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
func FindCharByCode(codes string, pointer int64, n *Node) (string, int64) {
	code := string(codes[pointer])
	if code == leftPathValue {
		if n.Left != nil && len(n.Left.Charachter) > 0 {
			return n.Left.Charachter, pointer
		}
		pointer++
		return FindCharByCode(codes, pointer, n.Left)
	} else if code == rightPathValue {
		if n.Right != nil && len(n.Right.Charachter) > 0 {
			return n.Right.Charachter, pointer
		}

		pointer++
		return FindCharByCode(codes, pointer, n.Right)
	}
	panic(fmt.Sprintf("Code %s is not supported", code))
}

//Node ...
type Node struct {
	Charachter string
	Weight     int64
	Left       *Node
	Right      *Node
}
