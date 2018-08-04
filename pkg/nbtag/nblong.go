package nbtag

import (
	"encoding/binary"
	"fmt"
)

type NBLong struct {
	tagData
	value int // TODO - does this need to be int64?
}

func parseLongTag(data []byte, pos int) (*NBLong, int) {
	tag := new(NBLong)
	tag.startPos = pos - 1
	tag.kind = NBTypeLong

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt64(data, pos)

	fmt.Printf("-> NBLong, name='%s', value='%d'\n", tag.name, tag.value)

	return tag, pos
}

func parseInt64(data []byte, pos int) (int, int) {
	val := int(binary.BigEndian.Uint64(data[pos : pos+8]))
	pos += 8

	return val, pos
}
