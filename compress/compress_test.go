package compress

import (
	"testing"
)

func TestCompress(t *testing.T) {

	t.Run("Compress - should pass", func(t *testing.T) {
		_, err := Compress("testdata/test_file_1.txt")
		if err != nil {
			t.Error(err)
		}
	})
}
