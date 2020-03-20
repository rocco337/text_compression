package compress

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
	"sort"
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

	file, err := os.Open(filePath)
	check(err)

	r := bufio.NewReader(file)
	tokens := make(map[string][]int)
	token := ""
	tokenPosition := 0

	for {
		c, _, err := r.ReadRune()
		if isEndOfFile(err) {
			break
		}
		check(err)

		token = token + string(c)
		if len(token) == defaultTokenSize {
			tokens[token] = append(tokens[token], tokenPosition)
			tokenPosition++
			token = ""
		}
	}

	outfile := filePath[0:len(filePath)-len(extension)] + outputExtension

	json, err := json.Marshal(tokens)
	b := new(bytes.Buffer)

	e := gob.NewEncoder(b)

	err = e.Encode(tokens)
	check(err)

	err = ioutil.WriteFile(outfile, b.Bytes(), 0644)

	outfile = filePath[0:len(filePath)-len(extension)] + ".json"
	err = ioutil.WriteFile(outfile, json, 0644)

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
