package nbtag

import (
	"encoding/binary"
	"math"
)

type NBFloat struct {
	tagData
	value float32
}

func parseFloatTag(data []byte, pos int) (*NBFloat, int) {
	tag := new(NBFloat)
	tag.startPos = pos - 1
	tag.kind = NBTypeFloat

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseFloat(data, pos)

	tagLog("-> NBFloat, name='%s', value=%.4f\n", tag.name, tag.value)

	return tag, pos
}

func parseFloatListItem(data []byte, pos int, name string) (*NBFloat, int) {
	tag := new(NBFloat)
	tag.startPos = pos
	tag.kind = NBTypeFloat
	tag.name = name

	tag.value, pos = parseFloat(data, pos)

	tagLog("-> NBFloat list item, name='%s', value=%.4f\n", tag.name, tag.value)

	return tag, pos
}

func parseFloat(data []byte, pos int) (float32, int) {
	bits := binary.BigEndian.Uint32(data[pos : pos+4])
	float := math.Float32frombits(bits)
	pos += 4
	return float, pos
}
