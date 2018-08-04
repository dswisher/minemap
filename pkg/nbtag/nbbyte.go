package nbtag

import "fmt"

type NBByte struct {
	kind  byte
	name  string
	value int
}

func parseByteTag(data []byte, pos int) (*NBByte, int) {
	tag := NBByte{kind: NBTypeByte}

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt8(data, pos)

	fmt.Printf("-> NBByte, name='%s', value='%d'\n", tag.name, tag.value)

	return &tag, pos
}

func (c *NBByte) GetType() byte {
	return c.kind
}

func (c *NBByte) GetName() string {
	return c.name
}
