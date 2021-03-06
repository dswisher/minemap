package nbtag

import "fmt"

type NBEnd struct {
	tagData
}

func newEndTag() *NBEnd {
	tag := new(NBEnd)
	tag.kind = NBTypeEnd

	return tag
}

// Parse an end tag, which does not have a name.
// The current position should be the byte following the tag type byte.
func (tag *NBEnd) Parse(reader NBReader) error {
	return nil
}

func (tag *NBEnd) String() string {
	return fmt.Sprintf("NBEnd: startPos=0x%04X", tag.startPos)
}
