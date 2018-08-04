package nbtag

import "fmt"

type NBLong struct {
	kind  byte
	name  string
	value int // TODO - does this need to be int64?
}

func parseLongTag(data []byte, pos int) (*NBLong, int) {
	tag := NBLong{kind: NBTypeLong}

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt64(data, pos)

	fmt.Printf("-> NBLong, name='%s', value='%d'\n", tag.name, tag.value)

	return &tag, pos
}

func (c *NBLong) GetType() byte {
	return c.kind
}

func (c *NBLong) GetName() string {
	return c.name
}
