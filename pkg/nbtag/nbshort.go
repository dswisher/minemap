package nbtag

import (
	"fmt"
)

type NBShort struct {
	tagData
	value int
}

func newShortTag() *NBShort {
	tag := new(NBShort)
	tag.kind = NBTypeShort

	return tag
}

// Parse a short tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBShort) Parse(reader NBReader) error {
	var err error
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	return tag.parseData(reader)
}

// Parse the data for a short.
func (tag *NBShort) parseData(reader NBReader) error {
	var err error
	tag.value, err = reader.ReadInt16()

	return err
}

func (tag *NBShort) String() string {
	return fmt.Sprintf("NBShort: startPos=0x%04X, val=%d, name='%s'", tag.startPos, tag.value, tag.name)
}
