package nbtag

import (
	"fmt"
)

type NBInt struct {
	tagData
	value int
}

func newIntTag() *NBInt {
	tag := new(NBInt)
	tag.kind = NBTypeInt

	return tag
}

// Parse an int tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBInt) Parse(reader NBReader) error {
	var err error
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	tag.value, err = reader.ReadInt32()

	return err
}

func (tag *NBInt) String() string {
	return fmt.Sprintf("NBInt: startPos=0x%04X, val=%d, name='%s'", tag.startPos, tag.value, tag.name)
}
