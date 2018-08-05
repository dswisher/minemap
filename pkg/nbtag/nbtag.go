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
	NBTypeFloat     = 5
	NBTypeDouble    = 6
	NBTypeByteArray = 7
	NBTypeString    = 8
	NBTypeList      = 9
	NBTypeCompound  = 10
	NBTypeIntArray  = 11
	NBTypeLongArray = 12
)

type NBTag interface {
	GetType() byte
	GetName() string
	GetStartPos() int
}

type tagData struct {
	startPos int
	kind     byte
	name     string
}

func Parse(data []byte, pos int) NBTag {
	tag, _ := parseTag(data, pos)

	return tag
}

func parseTag(data []byte, pos int) (NBTag, int) {
	var tag NBTag

	kind := data[pos]
	pos += 1

	switch kind {
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
	case NBTypeFloat:
		tag, pos = parseFloatTag(data, pos)
	case NBTypeDouble:
		tag, pos = parseDoubleTag(data, pos)
	case NBTypeByteArray:
		tag, pos = parseByteArrayTag(data, pos)
	case NBTypeString:
		tag, pos = parseStringTag(data, pos)
	case NBTypeList:
		tag, pos = parseListTag(data, pos)
	case NBTypeCompound:
		tag, pos = parseCompoundTag(data, pos)
	case NBTypeIntArray:
		tag, pos = parseIntArrayTag(data, pos)
	case NBTypeLongArray:
		tag, pos = parseLongArrayTag(data, pos)
	default:
		log.Fatalf("-> Unhandled NBT tag type %d at pos 0x%x.\n", kind, pos-1)
	}

	return tag, pos
}

func tagLog(format string, args ...interface{}) {
	// fmt.Printf(format, args...)
}

func (t *tagData) GetType() byte {
	return t.kind
}

func (t *tagData) GetName() string {
	return t.name
}

func (t *tagData) GetStartPos() int {
	return t.startPos
}
