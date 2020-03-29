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

	readerIndex := 0

	//read how long is huffmanTree in bytes
	huffmanTreeLengthBytes := make([]byte, 8)
	for readerIndex < len(huffmanTreeLengthBytes) {
		currentByte, err := reader.ReadByte()
		check(err)

		huffmanTreeLengthBytes[readerIndex] = currentByte
		readerIndex++
	}

	//read bytes of huffman tree
	huffmanTreeBytes := make([]byte, binary.LittleEndian.Uint64(huffmanTreeLengthBytes))
	i := 0
	for i < len(huffmanTreeBytes) {
		currentByte, err := reader.ReadByte()
		check(err)
		huffmanTreeBytes[i] = currentByte
		i++
		readerIndex++
	}

	//decode huffmanTree from bytes
	var huffmanTree huffman.Node
	dec := gob.NewDecoder(bytes.NewReader(huffmanTreeBytes))
	err = dec.Decode(&huffmanTree)
	check(err)

	i = 0
	//read total bit length of compressed content
	compressedContentBitsLength := make([]byte, 4)
	for i < 4 {
		currentByte, err := reader.ReadByte()
		check(err)

		compressedContentBitsLength[i] = currentByte
		readerIndex++
		i++
	}

	compressedContentTotalLength := binary.LittleEndian.Uint32(compressedContentBitsLength)

	//read compressed content
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

	decompressed := huffman.Decompress(compressedContentBytes, compressedContentTotalLength, &huffmanTree)
	return decompressed
}
