package nbtag

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
	// TODO - is this all we need?
	return nil
}

func parseEndTag(data []byte, pos int) (*NBEnd, int) {
	tag := new(NBEnd)
	tag.startPos = pos - 1
	tag.kind = NBTypeEnd
	tag.name = ""

	tagLog("-> NBEnd\n")

	return tag, pos
}
