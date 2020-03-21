package huffman

import (
	"testing"
)

func TestHuffman(t *testing.T) {

	t.Run("Compress - should pass", func(t *testing.T) {
		expected := "100110010111110100"
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

	t.Run("createHuffmanTree - should pass", func(t *testing.T) {
		characters := make(map[string]int64, 0)
		characters["f"] = 2
		characters["H"] = 1
		characters["a"] = 1
		characters["m"] = 1
		characters["u"] = 1
		characters["n"] = 1

		tree := createHuffmanTree(characters)

		if tree.Weight != 7 {
			t.Error("Expected: 7", "Actual :", tree.Weight)
		}

		if tree.Left.Weight != 3 {
			t.Error("Expected: 3", "Actual :", tree.Left.Weight)
		}

		if tree.Right.Weight != 4 {
			t.Error("Expected: 3", "Actual :", tree.Right.Weight)
		}
	})

}

func BenchmarkHuffmanCompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compress("Huffman")
	}
}
