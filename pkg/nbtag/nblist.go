package nbtag

import (
	"fmt"
	"io"
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
	tag.items = make([]NBTag, tag.count)

	for i := 0; i < tag.count; i++ {
		inner, err := newTag(reader, tag.innerType)
		if err != nil {
			return err
		}

		inner.SetName(fmt.Sprintf("%s[%d]", tag.Name(), i))

		tag.items[i] = inner

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
	return fmt.Sprintf("NBList: startPos=0x%04X, innerType=%d, count=%d, len(items)=%d, name='%s'",
		tag.startPos, tag.innerType, tag.count, len(tag.items), tag.name)
}

func (tag *NBList) DumpIndented(w io.Writer, depth int) {
	writeIndented(w, depth, tag.String())
	for i := 0; i < tag.count; i++ {
		tag.items[i].DumpIndented(w, depth+1)
	}
}
