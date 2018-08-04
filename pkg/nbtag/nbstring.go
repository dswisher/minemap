package nbtag

import "fmt"

type NBString struct {
	kind  byte
	name  string
	value string
}

func parseStringTag(data []byte, pos int) (*NBString, int) {
	tag := NBString{kind: NBTypeString}

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseString(data, pos)

	fmt.Printf("-> NBString, name='%s', value='%s'\n", tag.name, tag.value)

	return &tag, pos
}

func (c *NBString) GetType() byte {
	return c.kind
}

func (c *NBString) GetName() string {
	return c.name
}
