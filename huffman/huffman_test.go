package huffman

import (
	"testing"
)

func TestHuffman(t *testing.T) {

	t.Run("Compress - should pass", func(t *testing.T) {
		expected := "1000111110011101"

		actual, _ := Compress("Huffman")

		if expected != actual {
			t.Error("Expected " + expected + ", actual " + actual)
		}
	})

	// t.Run("Decompress - should pass", func(t *testing.T) {

	// 	codes, huffmanTree := Compress("Huffman")

	// 	Decompress(codes, huffmanTree)
	// })
}
