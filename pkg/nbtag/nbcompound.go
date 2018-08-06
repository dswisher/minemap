package nbtag

type NBCompound struct {
	tagData
	children map[string]NBTag
}

func (c *NBCompound) AddChildren(children []NBTag) {
	for _, child := range children {
		c.AddChild(child)
	}
}

func (c *NBCompound) AddChild(child NBTag) {
	c.children[child.GetName()] = child
}

func (c *NBCompound) GetChild(name string) NBTag {
	child, found := c.children[name]

	// Could probably just return child, but not sure what zero-value for NBTag is...
	if !found {
		return nil
	}

	return child
}

func newCompound() *NBCompound {
	tag := new(NBCompound)
	tag.kind = NBTypeCompound
	tag.children = make(map[string]NBTag)

	return tag
}

func parseCompoundTag(data []byte, pos int) (*NBCompound, int) {
	tag := newCompound()
	tag.startPos = pos - 1

	tag.name, pos = parseString(data, pos)

	tagLog("-> NBCompound, name='%s'\n", tag.name)

	children, pos := parseCompoundData(data, pos)

	tag.AddChildren(children)

	return tag, pos
}

func parseCompoundListItem(data []byte, pos int, name string) (*NBCompound, int) {
	tag := newCompound()
	tag.startPos = pos
	tag.name = name

	tagLog("-> NBCompound list item, name='%s'\n", tag.name)

	children, pos := parseCompoundData(data, pos)

	tag.AddChildren(children)

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
