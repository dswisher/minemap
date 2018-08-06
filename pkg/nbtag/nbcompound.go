package nbtag

type NBCompound struct {
	tagData
	// TODO - this should be a MAP!
	Children []NBTag
}

func parseCompoundTag(data []byte, pos int) (*NBCompound, int) {
	tag := new(NBCompound)
	tag.startPos = pos - 1
	tag.kind = NBTypeCompound

	tag.name, pos = parseString(data, pos)

	tagLog("-> NBCompound, name='%s'\n", tag.name)

	tag.Children, pos = parseCompoundData(data, pos)

	return tag, pos
}

func parseCompoundListItem(data []byte, pos int, name string) (*NBCompound, int) {
	tag := new(NBCompound)
	tag.startPos = pos
	tag.kind = NBTypeCompound
	tag.name = name

	tagLog("-> NBCompound list item, name='%s'\n", tag.name)

	tag.Children, pos = parseCompoundData(data, pos)

	return tag, pos
}

func parseCompoundData(data []byte, pos int) ([]NBTag, int) {
	var children []NBTag
	var child NBTag
	for ok := true; ok; ok = (child.GetType() != NBTypeEnd) {
		child, pos = parseTag(data, pos)
		children = append(children, child)
	}

	return children, pos
}
