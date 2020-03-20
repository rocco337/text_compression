package compress

import (
	"testing"
)

func TestCompress(t *testing.T) {

	t.Run("Compress - should pass", func(t *testing.T) {
		Compress("aasssssa")
	})
}
