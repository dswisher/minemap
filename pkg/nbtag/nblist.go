package nbtag

import (
	"fmt"
)

type NBList struct {
	tagData
	innerType byte
	count     int
	items     []NBTag
}

func newListTag() *NBList {
	tag := new(NBList)
	tag.kind = NBTypeList

	return tag
}

// Parse a list tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBList) Parse(reader NBReader) error {
	var err error

	// Name
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	return tag.parseData(reader)
}

func (tag *NBList) parseData(reader NBReader) error {
	var err error

	// Inner type
	tag.innerType, err = reader.ReadByte()
	if err != nil {
		return err
	}

	// Count
	tag.count, err = reader.ReadInt32()
	if err != nil {
		return err
	}

	// Parse the items
	for i := 0; i < tag.count; i++ {
		inner, err := newTag(reader, tag.innerType)
		if err != nil {
			return err
		}

		reader.PushContext(inner)
		inner.SetStartPos(reader.Pos())
		err = inner.parseData(reader)
		if err != nil {
			return err
		}

		reader.PopContext()
	}

	return nil
}

func (tag *NBList) String() string {
	return fmt.Sprintf("NBList: startPos=0x%04X, innerType=%d, count=%d, name='%s'",
		tag.startPos, tag.innerType, tag.count, tag.name)
}
