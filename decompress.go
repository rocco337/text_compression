package textcompression

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
	"os"
	"text_compression/huffman"
)

//Decompress ...
func Decompress(filePath string) string {
	file, err := os.Open(filePath)
	defer file.Close()
	check(err)

	reader := bufio.NewReader(file)

	i := 0
	huffmanTreeLengthBytes := make([]byte, 8)

	for i < 8 {
		currentByte, err := reader.ReadByte()
		check(err)

		huffmanTreeLengthBytes[i] = currentByte
		i++
	}
	huffmanTreeLength := binary.LittleEndian.Uint64(huffmanTreeLengthBytes)
	huffmanTreeBytes := make([]byte, huffmanTreeLength)
	j := 0
	for uint64(i) < huffmanTreeLength+8 {
		currentByte, err := reader.ReadByte()
		check(err)
		huffmanTreeBytes[j] = currentByte
		j++
		i++
	}

	var huffmanTree huffman.Node
	dec := gob.NewDecoder(bytes.NewReader(huffmanTreeBytes))
	err = dec.Decode(&huffmanTree)
	check(err)

	compressedContentBytes := make([]byte, 0)

	for {
		currentByte, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		compressedContentBytes = append(compressedContentBytes, currentByte)
	}

	decompressed := huffman.Decompress(compressedContentBytes, &huffmanTree)
	return decompressed
}
