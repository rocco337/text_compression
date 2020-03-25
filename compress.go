package textcompression

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"

	"text_compression/huffman"
)

var outputExtension = ".ztxt"
var supportedExtensions = []string{".txt"}
var defaultTokenSize = 3

//Compress ... reads a file, compress it and save new one
func Compress(filePath string, outputfilePath string) error {

	extension := path.Ext(filePath)
	if !contains(supportedExtensions, extension) {
		return errors.New("Unsupported extension: " + extension)
	}

	fileBytes, err := ioutil.ReadFile(filePath)
	check(err)

	compressed, huffman := huffman.Compress(fileBytes)

	huffmanTreeBytes, err := json.Marshal(huffman)
	check(err)

	file, err := os.Create(outputfilePath)
	defer file.Close()
	check(err)

	file.Write(huffmanTreeBytes)
	file.Write([]byte(fmt.Sprintln()))
	file.Write([]byte(compressed))

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
