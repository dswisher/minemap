package nbtag

import (
	"fmt"
)

type NBCompound struct {
	tagData
	children map[string]NBTag
}

// TODO - is this still used?
func (c *NBCompound) AddChildren(children []NBTag) {
	for _, child := range children {
		c.AddChild(child)
	}
}

func (c *NBCompound) AddChild(child NBTag) {
	c.children[child.GetName()] = child
}

func (c *NBCompound) ContainsChild(name string) bool {
	_, found := c.children[name]
	return found
}

func (c *NBCompound) GetChild(name string) NBTag {
	child, found := c.children[name]

	// Could probably just return child, but not sure what zero-value for NBTag is...
	if !found {
		return nil
	}

	return child
}

func newCompoundTag() *NBCompound {
	tag := new(NBCompound)
	tag.kind = NBTypeCompound
	tag.children = make(map[string]NBTag)

	return tag
}

// Parse a compound tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBCompound) Parse(reader NBReader) error {
	var err error
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	return tag.parseData(reader)
}

// Parse the data for a compound.
func (tag *NBCompound) parseData(reader NBReader) error {
	var child NBTag
	var err error
	for ok := true; ok; ok = (child.GetType() != NBTypeEnd) {
		child, err = parseTag(reader)
		if err != nil {
			return err
		}
		tag.AddChild(child)
	}

	return nil
}

func (tag *NBCompound) String() string {
	return fmt.Sprintf("NBCompound: startPos=0x%04X, len(children)=%d, name='%s'", tag.startPos, len(tag.children), tag.name)
}

// TODO - DELETEME
func parseCompoundTag(data []byte, pos int) (*NBCompound, int) {
	tag := newCompoundTag()
	tag.startPos = pos - 1

	tag.name, pos = parseString(data, pos)

	tagLog("-> NBCompound, name='%s'\n", tag.name)

	children, pos := parseCompoundData(data, pos)

	tag.AddChildren(children)

	return tag, pos
}

func parseCompoundListItem(data []byte, pos int, name string) (*NBCompound, int) {
	tag := newCompoundTag()
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
		child, pos = parseTagOld(data, pos)
		children = append(children, child)
	}

	return children, pos
}
