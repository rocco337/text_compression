package huffman

import (
	"container/heap"
	"errors"
	"math/big"
)

//Compress ...
func Compress(input []byte) ([]byte, uint, *Node) {
	//create all leaf nodes
	characters := createLeafNodes(input)

	//create tree from leaf nodes
	huffmanTree := createHuffmanTree(characters)

	output := big.Int{}
	outputByteIndex := 0
	var totalLength uint

	for _, b := range input {
		char := string(b)

		//find bits from code table
		codePath, err := FindCodePath(huffmanTree, char)
		check(err)
		//we have to hold length of all bits that are written
		totalLength += codePath.Len

		//write codePath bits to array tha holds output
		for _, bit := range codePath.Value {
			output.SetBit(&output, outputByteIndex, bit)
			outputByteIndex++
		}
	}

	return output.Bytes(), totalLength, huffmanTree
}

//Decompress ...
func Decompress(compressed []byte, totalLength uint, huffmanTree *Node) string {
	var decompressed = ""

	var byteArray big.Int
	byteArray.SetBytes(compressed)

	temp := BitArray{}
	var i uint

	for i < totalLength {
		temp.Append(byteArray.Bit(int(i)))

		char, err := FindCharacterByCodePath(huffmanTree, &temp, 0)
		if err == nil {
			decompressed += char
			temp = BitArray{}
		}

		i++
	}

	return decompressed
}

func createLeafNodes(input []byte) *PriorityQueue {
	nodes := make(PriorityQueue, 0)
	characters := make(map[string]*Node)

	i := 0
	for _, b := range input {
		char := string(b)
		if node, ok := characters[char]; ok {
			node.Weight++
		} else {
			node := &Node{}
			node.Charachter = char
			node.Weight = 1
			node.CodePath = new(BitArray)

			nodes = append(nodes, &Item{
				value: node,
				index: i,
			})
			characters[char] = node
		}

		i++
	}

	return &nodes
}

func createHuffmanTree(nodes *PriorityQueue) *Node {
	//initiate priority queue. Makes sure that we always pop nodes with lowest weight
	heap.Init(nodes)

	leftCodePath := uint(0)
	rightCodePath := uint(1)

	//construct huffman tree
	for nodes.Len() > 1 {
		newNode := &Node{}
		newNode.Nodes = make([]*Node, 2)
		newNode.CodePath = new(BitArray)

		newNode.Nodes[0] = nodes.Pop().(*Node)
		newNode.Nodes[1] = nodes.Pop().(*Node)

		newNode.Nodes[0].UpdateCodePath(leftCodePath)
		newNode.Nodes[1].UpdateCodePath(rightCodePath)

		newNode.Weight += newNode.Nodes[0].Weight + newNode.Nodes[1].Weight

		newPriorityQueueItem := &Item{
			value: newNode,
		}

		heap.Push(nodes, newPriorityQueueItem)
		nodes.update(newPriorityQueueItem, newPriorityQueueItem.value)
	}

	return (*nodes)[0].value
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

//FindCodePath ...
func FindCodePath(n *Node, char string) (*BitArray, error) {
	returnError := func() error {
		return errors.New("Cannot find character:" + char)
	}

	if n == nil {
		return nil, returnError()
	}

	if n.Charachter == char {
		return n.CodePath, nil
	}

	if len(n.Nodes) == 0 {
		return nil, returnError()
	}

	codePath, err := FindCodePath(n.Nodes[0], char)
	if err != nil {
		codePath, err = FindCodePath(n.Nodes[1], char)
		if err != nil {
			return nil, returnError()
		}

	}

	return codePath, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
