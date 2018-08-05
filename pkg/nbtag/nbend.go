package nbtag

type NBEnd struct {
	tagData
}

func parseEndTag(data []byte, pos int) (*NBEnd, int) {
	tag := new(NBEnd)
	tag.startPos = pos - 1
	tag.kind = NBTypeEnd
	tag.name = ""

	tagLog("-> NBEnd\n")

	return tag, pos
}
