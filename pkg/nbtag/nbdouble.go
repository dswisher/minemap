package nbtag

import (
	"encoding/binary"
	"fmt"
	"math"
)

type NBDouble struct {
	tagData
	value float64
}

func newDoubleTag() *NBDouble {
	tag := new(NBDouble)
	tag.kind = NBTypeDouble

	return tag
}

// Parse a double tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBDouble) Parse(reader NBReader) error {
	var err error
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	tag.value, err = reader.ReadDouble()

	return err
}

func (tag *NBDouble) String() string {
	return fmt.Sprintf("NBDouble: startPos=0x%04X, val=%.4f, name='%s'", tag.startPos, tag.value, tag.name)
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
