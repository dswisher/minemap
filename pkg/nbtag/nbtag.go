package nbtag

import (
	"encoding/binary"
	"log"
	"math"
)

const (
	NBTypeEnd       = 0
	NBTypeByte      = 1
	NBTypeShort     = 2
	NBTypeInt       = 3
	NBTypeLong      = 4
	NBTypeDouble    = 6
	NBTypeString    = 8
	NBTypeList      = 9
	NBTypeCompound  = 10
	NBTypeIntArray  = 11
	NBTypeLongArray = 12
)

type NBTag interface {
	GetType() byte
	GetName() string
}

func Parse(data []byte, pos int) NBTag {
	tag, _ := parseTag(data, pos)

	return tag
}

func parseTag(data []byte, pos int) (NBTag, int) {
	var tag NBTag

	t := data[pos]
	pos += 1

	switch t {
	case NBTypeEnd:
		tag, pos = parseEndTag(data, pos)
	case NBTypeByte:
		tag, pos = parseByteTag(data, pos)
	case NBTypeShort:
		tag, pos = parseShortTag(data, pos)
	case NBTypeInt:
		tag, pos = parseIntTag(data, pos)
	case NBTypeLong:
		tag, pos = parseLongTag(data, pos)
	case NBTypeDouble:
		tag, pos = parseDoubleTag(data, pos)
	case NBTypeCompound:
		tag, pos = parseCompoundTag(data, pos)
	case NBTypeList:
		tag, pos = parseListTag(data, pos)
	case NBTypeString:
		tag, pos = parseStringTag(data, pos)
	case NBTypeIntArray:
		tag, pos = parseIntArrayTag(data, pos)
	case NBTypeLongArray:
		tag, pos = parseLongArrayTag(data, pos)
	default:
		log.Fatalf("-> Unhandled NBT tag type %d at pos 0x%x.\n", t, pos-1)
	}

	return tag, pos
}

func parseString(data []byte, pos int) (string, int) {
	nameLen, pos := parseInt16(data, pos)

	var name string
	if nameLen > 0 {
		name = string(data[pos : pos+nameLen])
	} else {
		name = ""
	}

	pos += nameLen

	return name, pos
}

func parseDouble(data []byte, pos int) (float64, int) {
	bits := binary.BigEndian.Uint64(data[pos : pos+8])
	float := math.Float64frombits(bits)
	pos += 8
	return float, pos
}

func parseInt64(data []byte, pos int) (int, int) {
	val := int(binary.BigEndian.Uint64(data[pos : pos+8]))
	pos += 8

	return val, pos
}

func parseInt32(data []byte, pos int) (int, int) {
	val := int(binary.BigEndian.Uint32(data[pos : pos+4]))
	pos += 4

	return val, pos
}

func parseInt16(data []byte, pos int) (int, int) {
	val := int(binary.BigEndian.Uint16(data[pos : pos+2]))
	pos += 2

	return val, pos
}

func parseInt8(data []byte, pos int) (int, int) {
	val := int(data[pos])
	pos += 1

	return val, pos
}
