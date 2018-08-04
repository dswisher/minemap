package nbtag

import "fmt"

type NBDouble struct {
	kind  byte
	name  string
	value float64
}

func parseDoubleTag(data []byte, pos int) (*NBDouble, int) {
	tag := NBDouble{kind: NBTypeDouble}

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseDouble(data, pos)

	fmt.Printf("-> NBDouble, name='%s', value='%.4f'\n", tag.name, tag.value)

	return &tag, pos
}

func (c *NBDouble) GetType() byte {
	return c.kind
}

func (c *NBDouble) GetName() string {
	return c.name
}
