package nbtag

import (
	"encoding/binary"
	"fmt"
)

type NBInt struct {
	tagData
	value int
}

func parseIntTag(data []byte, pos int) (*NBInt, int) {
	tag := new(NBInt)
	tag.kind = NBTypeInt

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt32(data, pos)

	fmt.Printf("-> NBInt, name='%s', value='%d'\n", tag.name, tag.value)

	return tag, pos
}

func parseInt32(data []byte, pos int) (int, int) {
	val := int(binary.BigEndian.Uint32(data[pos : pos+4]))
	pos += 4

	return val, pos
}
