package nbtag

import (
	"encoding/binary"
	"math"
)

type NBDouble struct {
	tagData
	value float64
}

func parseDoubleTag(data []byte, pos int) (*NBDouble, int) {
	tag := new(NBDouble)
	tag.startPos = pos - 1
	tag.kind = NBTypeDouble

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseDouble(data, pos)

	tagLog("-> NBDouble, name='%s', value=%.4f\n", tag.name, tag.value)

	return tag, pos
}

func parseDoubleListItem(data []byte, pos int, name string) (*NBDouble, int) {
	tag := new(NBDouble)
	tag.startPos = pos
	tag.kind = NBTypeDouble
	tag.name = name

	tag.value, pos = parseDouble(data, pos)

	tagLog("-> NBDouble list item, name='%s', value=%.4f\n", tag.name, tag.value)

	return tag, pos
}

func parseDouble(data []byte, pos int) (float64, int) {
	bits := binary.BigEndian.Uint64(data[pos : pos+8])
	float := math.Float64frombits(bits)
	pos += 8
	return float, pos
}
