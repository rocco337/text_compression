package huffman

import (
	"math/big"
	"testing"
)

func TestHuffman(t *testing.T) {

	t.Run("Compress - should pass", func(t *testing.T) {
		expected := []uint{1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 0, 0, 1, 1}
		compressedBytes, _, _ := Compress([]byte("Huffman"))

		var compressedBits big.Int

		compressedBits.SetBytes(compressedBytes)

		i := 0
		for i < compressedBits.BitLen() {
			if compressedBits.Bit(i) != expected[i] {
				t.Error("Error at index", i, "Expected: ", expected[i], "Actual: ", compressedBits.Bit(i))
				break
			}
			i++
		}
	})

	t.Run("Decompress - should pass", func(t *testing.T) {
		expected := "Huffman"
		expectedBytes := []uint{1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 0, 0, 1, 1}

		compressedBytes, totalLength, huffmanTree := Compress([]byte(expected))

		var compressedBits big.Int
		compressedBits.SetBytes(compressedBytes)

		i := 0
		for i < compressedBits.BitLen() {
			if compressedBits.Bit(i) != expectedBytes[i] {
				t.Error("Error at index", i, "Expected: ", expectedBytes[i], "Actual: ", compressedBits.Bit(i))
				break
			}
			i++
		}

		actual := Decompress(compressedBytes, totalLength, huffmanTree)
		if expected != actual {
			t.Error("Expected " + expected + ", actual " + actual)
		}
	})

	// t.Run("createHuffmanTree - should pass", func(t *testing.T) {
	// 	leafNodes := make(PriorityQueue, 0)
	// 	leafNodes.Push(&Item{ &Node })
	// 	characters["f"] = 2
	// 	characters["a"] = 1
	// 	characters["m"] = 1
	// 	characters["u"] = 1
	// 	characters["n"] = 1
	// 	characters["H"] = 1

	// 	tree := createHuffmanTree(characters)

	// 	if tree.Weight != 7 {
	// 		t.Error("Expected: 7", "Actual :", tree.Weight)
	// 	}

	// 	if tree.Nodes[0].Weight != 3 {
	// 		t.Error("Expected: 3", "Actual :", tree.Nodes[0].Weight)
	// 	}

	// 	if tree.Nodes[1].Weight != 4 {
	// 		t.Error("Expected: 3", "Actual :", tree.Nodes[1].Weight)
	// 	}
	// })

	t.Run("Decompress - should pass", func(t *testing.T) {
		expected := "Huffman is good algorithm"
		compressed, totalLength, huffmanTree := Compress([]byte(expected))

		actual := Decompress(compressed, totalLength, huffmanTree)
		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})

}

func BenchmarkHuffmanCompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compress([]byte("Huffman"))
	}
}
