package compress

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"sort"

	"text_compression/huffman"
)

var outputExtension = ".ztxt"
var supportedExtensions = []string{".txt"}
var defaultTokenSize = 3

//Compress ... reads a file, compress it and save new one
func Compress(filePath string) (string, error) {

	extension := path.Ext(filePath)
	if !contains(supportedExtensions, extension) {
		return "", errors.New("Unsupported extension: " + extension)
	}

	fileBytes, err := ioutil.ReadFile(filePath)
	check(err)

	compressed, huffman := huffman.Compress(string(fileBytes))

	fmt.Println(len(compressed))
	js, _ := json.Marshal(huffman)

	outfile := filePath[0:len(filePath)-len(extension)] + outputExtension
	ioutil.WriteFile(outfile+".json", js, 0644)

	err = ioutil.WriteFile(outfile, compressed, 0644)

	check(err)

	return "", nil
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

func isEndOfFile(e error) bool {
	if e == nil {
		return false
	}

	return e == io.EOF
}
