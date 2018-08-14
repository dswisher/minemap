package nbtag

import (
	"encoding/binary"
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

	tag.value, err = reader.ReadInt16()

	return err
}

func (tag *NBShort) String() string {
	return fmt.Sprintf("NBShort: startPos=0x%04X, val=%d, name='%s'", tag.startPos, tag.value, tag.name)
}

func parseShortTag(data []byte, pos int) (*NBShort, int) {
	tag := new(NBShort)
	tag.startPos = pos - 1
	tag.kind = NBTypeShort

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt16(data, pos)

	tagLog("-> NBShort, name='%s', value='%d'\n", tag.name, tag.value)

	return tag, pos
}

func parseInt16(data []byte, pos int) (int, int) {
	val := int(binary.BigEndian.Uint16(data[pos : pos+2]))
	pos += 2

	return val, pos
}
