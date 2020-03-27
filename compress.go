package textcompression

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io/ioutil"
	"os"
	"sort"
	"strconv"

	"text_compression/huffman"
)

//Compress ... reads a file, compress it and save new one
func Compress(filePath string, outputfilePath string) error {
	fileBytes, err := ioutil.ReadFile(filePath)
	check(err)

	compressed, huffmanTree := huffman.Compress(fileBytes)

	var huffmanTreeBytes bytes.Buffer
	enc := gob.NewEncoder(&huffmanTreeBytes)
	err = enc.Encode(huffmanTree)
	check(err)

	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(huffmanTreeBytes.Len()))
	file, err := os.Create(outputfilePath)
	defer file.Close()
	check(err)

	file.Write(bs)
	file.Write(huffmanTreeBytes.Bytes())
	file.Write(compressed)

	check(err)

	return nil
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func asByteSlice(b string) []byte {
	var out []byte
	var str string

	for i := len(b); i > 0; i -= 8 {
		if i-8 < 0 {
			str = string(b[0:i])
		} else {
			str = string(b[i-8 : i])
		}
		v, err := strconv.ParseUint(str, 2, 8)
		if err != nil {
			panic(err)
		}
		out = append([]byte{byte(v)}, out...)
	}
	return out
}
