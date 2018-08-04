package nbtag

import "fmt"

type NBShort struct {
	kind  byte
	name  string
	value int
}

func parseShortTag(data []byte, pos int) (*NBShort, int) {
	tag := NBShort{kind: NBTypeShort}

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt16(data, pos)

	fmt.Printf("-> NBShort, name='%s', value='%d'\n", tag.name, tag.value)

	return &tag, pos
}

func (c *NBShort) GetType() byte {
	return c.kind
}

func (c *NBShort) GetName() string {
	return c.name
}
