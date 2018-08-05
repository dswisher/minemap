package nbtag

type NBCompound struct {
	tagData
	children []NBTag
}

func parseCompoundTag(data []byte, pos int) (*NBCompound, int) {
	tag := new(NBCompound)
	tag.startPos = pos - 1
	tag.kind = NBTypeCompound

	tag.name, pos = parseString(data, pos)

	tagLog("-> NBCompound, name='%s'\n", tag.name)

	tag.children, pos = parseCompoundData(data, pos)

	return tag, pos
}

func parseCompoundListItem(data []byte, pos int, name string) (*NBCompound, int) {
	tag := new(NBCompound)
	tag.startPos = pos
	tag.kind = NBTypeCompound
	tag.name = name

	tagLog("-> NBCompound list item, name='%s'\n", tag.name)

	tag.children, pos = parseCompoundData(data, pos)

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
