package huffman

import (
	"fmt"
	"strings"
)

//BitArray
type BitArray struct {
	Value []uint32
	Len   uint
}

// Preprend ....
func (b *BitArray) Preprend(val uint32) {
	b.Value = append(b.Value, val)
	b.Len++
}

// Preprend ....
func (b *BitArray) String() string {
	result := new(strings.Builder)

	var i int
	i = int(b.Len - 1)
	for i >= 0 {
		result.WriteString(fmt.Sprintf("%d", b.Value[i]))
		i--
	}

	return result.String()
}
