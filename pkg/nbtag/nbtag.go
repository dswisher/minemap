package nbtag

import (
	"log"
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

type tagData struct {
	kind byte
	name string
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

func (c *tagData) GetType() byte {
	return c.kind
}

func (c *tagData) GetName() string {
	return c.name
}
