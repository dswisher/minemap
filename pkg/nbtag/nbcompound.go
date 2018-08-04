package nbtag

import "fmt"

type NBCompound struct {
	kind     byte
	name     string
	children []NBTag
}

func parseCompoundTag(data []byte, pos int) (*NBCompound, int) {
	tag := NBCompound{kind: NBTypeCompound}

	tag.name, pos = parseString(data, pos)

	fmt.Printf("-> NBCompound, name='%s'\n", tag.name)

	tag.children, pos = parseCompoundData(data, pos)

	return &tag, pos
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

func (c *NBCompound) GetType() byte {
	return c.kind
}

func (c *NBCompound) GetName() string {
	return c.name
}
