package huffman

import (
	"container/heap"
	"errors"
	"math/big"
	"sort"
)

//Compress ...
func Compress(input []byte) ([]byte, *HuffmanTree) {
	characters := make(map[string]*Frequency)

	for _, b := range input {
		freq := new(Frequency)
		freq.Key = string(b)

		if characters[freq.Key] == nil {
			characters[freq.Key] = freq
		}

		characters[freq.Key].Count++
	}

	//create tree
	huffmanTree := createHuffmanTree(characters)

	codeTables := make(map[string]*BitArray)

	GenerateCodeTable(huffmanTree, &codeTables)

	//construct codes from tree
	temp := big.Int{}
	byteIndex := 0
	var totalLength uint

	for _, b := range input {
		char := string(b)
		bitArray := codeTables[char]
		totalLength += bitArray.Len
		for _, bit := range bitArray.Value {
			temp.SetBit(&temp, byteIndex, bit)
			byteIndex++
		}
	}

	return temp.Bytes(), &HuffmanTree{Root: huffmanTree, Length: totalLength}
}

//Decompress ...
func Decompress(compressed []byte, huffmanTree *HuffmanTree) string {
	var decompressed = ""

	var byteArray big.Int
	byteArray.SetBytes(compressed)

	temp := BitArray{}
	i := 0

	for uint(i) < huffmanTree.Length {
		temp.Append(byteArray.Bit(i))

		char, err := FindCharacterByCodePath(huffmanTree.Root, &temp, 0)
		if err == nil {
			decompressed += char
			temp = BitArray{}
		}

		i++
	}

	return decompressed
}

func createHuffmanTree(characters map[string]*Frequency) *Node {
	nodes := make(PriorityQueue, len(characters))

	i := 0

	//create all characters as leaf nodes. Make sure they are sorted
	for _, key := range getSortedMapKeys(characters) {
		node := &Node{}
		node.Charachter = key
		node.Weight = characters[key].Count
		node.CodePath = new(BitArray)

		nodes[i] = &Item{
			value: node,
			index: i,
		}
		i++
	}

	//initiate priority queue. Makes sure that we always pop nodes with lowest weight
	heap.Init(&nodes)

	leftCodePath := uint(0)
	rightCodePath := uint(1)

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

func getSortedMapKeys(characters map[string]*Frequency) []string {
	keys := make([]string, 0, len(characters))
	for k := range characters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

//GenerateCodeTable ...
func GenerateCodeTable(n *Node, codesTable *map[string]*BitArray) {
	if n == nil {
		return
	}

	if len(n.Charachter) > 0 {
		(*codesTable)[n.Charachter] = n.CodePath
	}

	for _, node := range n.Nodes {
		GenerateCodeTable(node, codesTable)
	}
}

//todo remove
//Frequency ...
type Frequency struct {
	Key      string
	Count    uint64
	CodePath []byte
}

//FindCharacterByCodePath ...
func FindCharacterByCodePath(n *Node, codePath *BitArray, pointer uint) (string, error) {
	if len(n.Charachter) > 0 && n.CodePath.Equals(codePath) {
		return n.Charachter, nil
	}

	if pointer > codePath.Len-1 {
		return "", errors.New("cannot find character")
	}

	lastBit := codePath.Value[pointer]
	pointer++

	if lastBit == 0 && len(n.Nodes) > 0 {
		return FindCharacterByCodePath(n.Nodes[0], codePath, pointer)
	}

	if lastBit == 1 && len(n.Nodes) > 1 {
		return FindCharacterByCodePath(n.Nodes[1], codePath, pointer)
	}

	return "", errors.New("Cannot find chracter")
}

//HuffmanTree ...
type HuffmanTree struct {
	Root   *Node
	Length uint
}
