package nbtag

import (
	"fmt"
)

type NBFloat struct {
	tagData
	value float32
}

func newFloatTag() *NBFloat {
	tag := new(NBFloat)
	tag.kind = NBTypeFloat

	return tag
}

// Parse a float tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBFloat) Parse(reader NBReader) error {
	var err error
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	return tag.parseData(reader)
}

// Parse the data for a float.
func (tag *NBFloat) parseData(reader NBReader) error {
	var err error
	tag.value, err = reader.ReadFloat()

	return err
}

func (tag *NBFloat) String() string {
	return fmt.Sprintf("NBFloat: startPos=0x%04X, val=%.4f, name='%s'", tag.startPos, tag.value, tag.name)
}
