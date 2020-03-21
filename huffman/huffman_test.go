package huffman

import (
	"testing"
)

func TestHuffman(t *testing.T) {

	t.Run("Compress - should pass", func(t *testing.T) {
		expected := "010001111100011101"
		actual, _ := Compress("Huffman")

		for i := range actual {
			if actual[i] != expected[i] {
				t.Error("Expected : ", expected, "Actual: ", actual)
				break

			}
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
