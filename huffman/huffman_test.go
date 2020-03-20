package huffman

import (
	"testing"
)

func TestHuffman(t *testing.T) {

	t.Run("Compress - should pass", func(t *testing.T) {
		expected := "010001111100011101"

		actual, _ := Compress("Huffman")

		if expected != actual {
			t.Error("Expected " + expected + ", actual " + actual)
		}
	})

	t.Run("Decompress - should pass", func(t *testing.T) {
		expected := "Huffman"
		codes, huffmanTree := Compress(expected)

		actual := Decompress(codes, huffmanTree)
		if expected != actual {
			t.Error("Expected " + expected + ", actual " + actual)
		}
	})
}
