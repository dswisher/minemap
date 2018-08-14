package nbtag

import (
	"fmt"
)

type NBDouble struct {
	tagData
	value float64
}

func newDoubleTag() *NBDouble {
	tag := new(NBDouble)
	tag.kind = NBTypeDouble

	return tag
}

// Parse a double tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBDouble) Parse(reader NBReader) error {
	var err error
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	return tag.parseData(reader)
}

// Parse the data for a double.
func (tag *NBDouble) parseData(reader NBReader) error {
	var err error
	tag.value, err = reader.ReadDouble()

	return err
}

func (tag *NBDouble) String() string {
	return fmt.Sprintf("NBDouble: startPos=0x%04X, val=%.4f, name='%s'", tag.startPos, tag.value, tag.name)
}
