package nbtag

import "fmt"

type NBInt struct {
	kind  byte
	name  string
	value int
}

func parseIntTag(data []byte, pos int) (*NBInt, int) {
	tag := NBInt{kind: NBTypeInt}

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt32(data, pos)

	fmt.Printf("-> NBInt, name='%s', value='%d'\n", tag.name, tag.value)

	return &tag, pos
}

func (c *NBInt) GetType() byte {
	return c.kind
}

func (c *NBInt) GetName() string {
	return c.name
}
