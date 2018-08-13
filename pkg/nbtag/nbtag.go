package nbtag

import (
	"fmt"
	"io"
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
	GetType() byte    // TODO - change to Type()
	GetName() string  // TODO - change to Name()
	GetStartPos() int // TODO - change to StartPos()

	Parse(reader NBReader) error

	Dump(w io.Writer)
}

type tagData struct {
	startPos int
	kind     byte
	name     string
}

// Main entry point for the named binary parser. Parse reads the
// byte stream and parses it into tags. It assumes the next byte
// to be read is the tag type.
func Parse(reader NBReader) (NBTag, error) {
	return parseTag(reader)
}

// The internal parse method that does all the real work. It is called
// internally when parsing things like an compound tag.
func parseTag(reader NBReader) (NBTag, error) {
	kind, err := reader.ReadByte()
	if err != nil {
		return nil, fmt.Errorf("parseTag: %s, pos 0x%X: %s", reader.Source(), reader.LastPos(), err)
	}

	var tag NBTag

	switch kind {
	case NBTypeCompound:
		tag = newCompound()
	default:
		return nil, fmt.Errorf("parseTag: %s, pos 0x%X: unhandled NBT tag type %d", reader.Source(), reader.LastPos(), kind)
	}

	err = tag.Parse(reader)
	if err != nil {
		return nil, err
	}

	return tag, nil
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

func (t *tagData) Dump(w io.Writer) {
	fmt.Fprintf(w, "Dump is not yet implemented!\n")
}

func (t *tagData) Parse(reader NBReader) error {
	return fmt.Errorf("Parse is not yet implemented for kind %d.", t.kind)
}

// - - - - OLD CODE BELOW, due to be deprecated

func ParseOld(data []byte, pos int) NBTag {
	tag, _ := parseTagOld(data, pos)

	return tag
}

func parseTagOld(data []byte, pos int) (NBTag, int) {
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
