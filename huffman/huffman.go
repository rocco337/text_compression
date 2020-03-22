package huffman

import (
	"container/heap"
	"errors"
	"sort"
)

var leftPathValue = "0"
var rightPathValue = "1"

//Compress ...
func Compress(text string) (string, *Node) {
	characters := make(map[string]int64)

	//go through text and get characters and frequencies
	for _, c := range text {
		char := string(c)
		characters[char]++
	}

	//create tree
	huffmanTree := createHuffmanTree(characters)

	codeTables := make(map[string]string)

	GenerateCodeTable(huffmanTree, &codeTables)

	//construct codes from tree
	compressed := ""
	for _, c := range text {
		char := string(c)
		compressed += codeTables[char]
	}

	return compressed, huffmanTree
}

//Decompress ...
func Decompress(compressed string, huffmanTree *Node) string {

	codeTables := make(map[string]string)
	GenerateCodeTable(huffmanTree, &codeTables)
	codeCharTable := InverseMap(&codeTables)

	var pointer int64
	var decompressed = ""
	codeBuffer := ""
	pointer = 0
	for pointer < int64(len(compressed)) {
		codeBuffer += string(compressed[pointer])
		if char, ok := codeCharTable[codeBuffer]; ok {
			decompressed += char
			codeBuffer = ""
		}
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
		newNode.Left.Move(leftPathValue)

		newNode.Right = nodes.Pop().(*Node)
		newNode.Right.Move(rightPathValue)

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
	for k := range characters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

//GetCodePathForChar ...
func GetCodePathForChar(n *Node, char string) (string, error) {
	if n == nil {
		return "", errors.New("Reached to end leaf without matching character")
	}

	if n.Left != nil && n.Left.Charachter == char {
		return n.Left.CodePath, nil
	} else if n.Right != nil && n.Right.Charachter == char {
		return n.Right.CodePath, nil
	} else {
		var err error

		code, err := GetCodePathForChar(n.Left, char)

		if err == nil {
			return code, nil
		}

		code, err = GetCodePathForChar(n.Right, char)
		if err != nil {
			return "", err
		}
		return code, nil
	}
}

//GenerateCodeTable ...
func GenerateCodeTable(n *Node, codesTable *map[string]string) {
	if n == nil {
		return
	}

	if len(n.Charachter) > 0 {
		(*codesTable)[n.Charachter] = n.CodePath
		return
	}

	if n.Left != nil {
		GenerateCodeTable(n.Left, codesTable)
	}

	if n.Right != nil {
		GenerateCodeTable(n.Right, codesTable)
	}
}

//InverseMap ...
func InverseMap(codes *map[string]string) map[string]string {
	inversed := make(map[string]string, 0)

	for key, value := range *codes {
		inversed[value] = key
	}

	return inversed
}

//Node ...
type Node struct {
	Charachter string
	Weight     int64
	CodePath   string
	Left       *Node
	Right      *Node
}

func (n *Node) Move(side string) {
	n.CodePath = side + n.CodePath

	if n.Left != nil {
		n.Left.Move(side)
	}

	if n.Right != nil {
		n.Right.Move(side)
	}
}
