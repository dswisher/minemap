package nbtag

import (
	"encoding/binary"
	"fmt"
)

type NBShort struct {
	tagData
	value int
}

func parseShortTag(data []byte, pos int) (*NBShort, int) {
	tag := new(NBShort)
	tag.kind = NBTypeShort

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt16(data, pos)

	fmt.Printf("-> NBShort, name='%s', value='%d'\n", tag.name, tag.value)

	return tag, pos
}

func parseInt16(data []byte, pos int) (int, int) {
	val := int(binary.BigEndian.Uint16(data[pos : pos+2]))
	pos += 2

	return val, pos
}
