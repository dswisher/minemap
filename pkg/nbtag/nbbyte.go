package nbtag

import "fmt"

type NBByte struct {
	tagData
	value int
}

func parseByteTag(data []byte, pos int) (*NBByte, int) {
	tag := new(NBByte)
	tag.startPos = pos - 1
	tag.kind = NBTypeByte

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt8(data, pos)

	fmt.Printf("-> NBByte, name='%s', value='%d'\n", tag.name, tag.value)

	return tag, pos
}

func parseInt8(data []byte, pos int) (int, int) {
	val := int(data[pos])
	pos += 1

	return val, pos
}
