package nbtag

import (
	"fmt"
)

type NBLong struct {
	tagData
	value int // TODO - does this need to be int64?
}

func newLongTag() *NBLong {
	tag := new(NBLong)
	tag.kind = NBTypeLong

	return tag
}

// Parse an long tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBLong) Parse(reader NBReader) error {
	var err error
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	tag.value, err = reader.ReadInt64()

	return err
}

func (tag *NBLong) String() string {
	return fmt.Sprintf("NBLong: startPos=0x%04X, val=%d, name='%s'", tag.startPos, tag.value, tag.name)
}
