package nbtag

import (
	"fmt"
	"io"
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
	fmt.Stringer

	GetType() byte    // TODO - change to Type()
	GetName() string  // TODO - change to Name()
	GetStartPos() int // TODO - change to StartPos()
	SetStartPos(pos int)

	Parse(reader NBReader) error
	parseData(reader NBReader) error

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

func newTag(reader NBReader, kind byte) (NBTag, error) {
	var tag NBTag

	switch kind {
	case NBTypeEnd:
		tag = newEndTag()
	case NBTypeByte:
		tag = newByteTag()
	case NBTypeShort:
		tag = newShortTag()
	case NBTypeInt:
		tag = newIntTag()
	case NBTypeLong:
		tag = newLongTag()
	case NBTypeFloat:
		tag = newFloatTag()
	case NBTypeDouble:
		tag = newDoubleTag()
	case NBTypeByteArray:
		tag = newByteArrayTag()
	case NBTypeString:
		tag = newStringTag()
	case NBTypeList:
		tag = newListTag()
	case NBTypeCompound:
		tag = newCompoundTag()
	case NBTypeIntArray:
		tag = newIntArrayTag()
	case NBTypeLongArray:
		tag = newLongArrayTag()
	default:
		return nil, newErrorf(reader, "Unhandled tag kind, %d, in newTag.", kind)
	}

	return tag, nil
}

// The internal parse method that does all the real work. It is called
// internally when parsing things like an compound tag.
func parseTag(reader NBReader) (NBTag, error) {
	kind, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	tag, err := newTag(reader, kind)
	if err != nil {
		return nil, err
	}

	reader.PushContext(tag)
	tag.SetStartPos(reader.Pos() - 1)
	err = tag.Parse(reader)
	if err != nil {
		return nil, err
	}

	reader.PopContext()

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

func (t *tagData) SetStartPos(pos int) {
	t.startPos = pos
}

func (t *tagData) Dump(w io.Writer) {
	fmt.Fprintf(w, "Dump is not yet implemented!\n")
}

func (t *tagData) parseData(reader NBReader) error {
	return newErrorf(reader, "parseData is not yet implemented for kind %d.", t.kind)
}

// - - - - OLD CODE BELOW, due to be deprecated

func tagLog(format string, args ...interface{}) {
	// fmt.Printf(format, args...)
}
