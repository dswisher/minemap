package nbtag

import (
	"encoding/binary"
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

func parseLongTag(data []byte, pos int) (*NBLong, int) {
	tag := new(NBLong)
	tag.startPos = pos - 1
	tag.kind = NBTypeLong

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt64(data, pos)

	tagLog("-> NBLong, name='%s', value='%d'\n", tag.name, tag.value)

	return tag, pos
}

func parseInt64(data []byte, pos int) (int, int) {
	val := int(binary.BigEndian.Uint64(data[pos : pos+8]))
	pos += 8

	return val, pos
}
