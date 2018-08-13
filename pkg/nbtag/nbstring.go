package nbtag

import "fmt"

type NBString struct {
	tagData
	value string
}

func newStringTag() *NBString {
	tag := new(NBString)
	tag.kind = NBTypeString

	return tag
}

// Parse a string tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBString) Parse(reader NBReader) error {
	var err error
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	// TODO - implement me - parse the data (the string itself)

	return newErrorf(reader, "NBString: Parse is not yet implemented")
}

func (tag *NBString) String() string {
	return fmt.Sprintf("NBString: startPos=0x%04X, name='%s', value='%s'", tag.startPos, tag.name, tag.value)
}

func parseStringTag(data []byte, pos int) (*NBString, int) {
	tag := new(NBString)
	tag.startPos = pos - 1
	tag.kind = NBTypeString

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseString(data, pos)

	tagLog("-> NBString, name='%s', value='%s'\n", tag.name, tag.value)

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
