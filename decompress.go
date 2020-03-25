package textcompression

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"text_compression/huffman"
)

//Decompress ...
func Decompress(filePath string) string {
	file, err := os.Open(filePath)
	defer file.Close()
	check(err)

	scanner := bufio.NewReader(file)
	huffmanTreeBytes, err := scanner.ReadBytes('\n')
	check(err)

	var huffmanTree huffman.HuffmanTree
	err = json.Unmarshal(huffmanTreeBytes, &huffmanTree)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	compressedContentBytes, err := scanner.ReadBytes('\n')
	if err != io.EOF {
		panic(err)
	}

	decompressed := huffman.Decompress(compressedContentBytes, &huffmanTree)
	return decompressed
}
