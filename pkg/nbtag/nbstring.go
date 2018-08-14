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

	tag.value, err = reader.ReadString()
	if err != nil {
		return err
	}

	return nil
}

func (tag *NBString) String() string {
	return fmt.Sprintf("NBString: startPos=0x%04X, name='%s', value='%s'", tag.startPos, tag.name, tag.value)
}
