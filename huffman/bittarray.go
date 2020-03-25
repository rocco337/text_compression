package huffman

import "errors"

//BitArray
type BitArray struct {
	Value []uint
	Len   uint
}

// Prepend ....
func (b *BitArray) Prepend(val uint) {
	b.Value = append([]uint{val}, b.Value...)
	b.Len++
}

// Append ....
func (b *BitArray) Append(val uint) {
	b.Value = append(b.Value, val)
	b.Len++
}

//Equals ...
func (b *BitArray) Equals(input *BitArray) bool {
	if b.Len != input.Len {
		return false
	}

	var i uint
	i = 0
	for i < b.Len {
		if b.Value[i] != input.Value[i] {
			return false
		}
		i++
	}

	return true
}

//GetLastBit ...
func (b *BitArray) GetLastBit() (uint, error) {
	if b.Len < 1 {
		return 0, errors.New("Empty array")
	}
	return b.Value[b.Len-1], nil
}
