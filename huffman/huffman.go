package huffman

import (
	"container/heap"
	"sort"
)

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

	// Conver map where key is character and value codepath to map where key is codemap and value is character
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

	//create all characters as leaf nodes. Make sure they are sorted
	for _, key := range getSortedMapKeys(characters) {
		node := &Node{}
		node.Charachter = key
		node.Weight = characters[key]
		node.CodePath = new(BitArray)

		nodes[i] = &Item{
			value: node,
			index: i,
		}
		i++
	}

	//initiate priority queue. Makes sure that we always pop nodes with lowest weight
	heap.Init(&nodes)

	leftCodePath := uint32(0)
	rightCodePath := uint32(1)

	//construct huffman tree
	for len(nodes) > 1 {
		newNode := &Node{}
		newNode.Nodes = make([]*Node, 2)
		newNode.CodePath = new(BitArray)

		newNode.Nodes[0] = nodes.Pop().(*Node)
		newNode.Nodes[1] = nodes.Pop().(*Node)

		newNode.Nodes[0].UpdateCodePath(leftCodePath)
		newNode.Nodes[1].UpdateCodePath(rightCodePath)

		for _, node := range newNode.Nodes {
			//sum weights
			newNode.Weight += node.Weight
		}

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

//GenerateCodeTable ...
func GenerateCodeTable(n *Node, codesTable *map[string]string) {
	if n == nil {
		return
	}

	if len(n.Charachter) > 0 {
		(*codesTable)[n.Charachter] = n.CodePath.String()
	}

	for _, node := range n.Nodes {
		GenerateCodeTable(node, codesTable)
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
