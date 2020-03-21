package textcompression

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestCompress(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	t.Run("Compress - Decompress - should pass", func(t *testing.T) {
		outFile := "testdata/test_file_" + fmt.Sprintf("%d", rand.Intn(1000)) + ".ztxt"
		inFile := "testdata/test_file_1.txt"

		err := Compress(inFile, outFile)
		if err != nil {
			t.Error(err)
		}

		decompressed := Decompress(outFile)

		originalFile, _ := ioutil.ReadFile(inFile)

		if string(originalFile) != decompressed {
			t.Error("Decompressed file content does not equal origina file")
		} else {
			os.Remove(outFile)
		}

	})
}
