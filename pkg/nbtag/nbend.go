package nbtag

import "fmt"

type NBEnd struct {
	kind byte
}

func parseEndTag(data []byte, pos int) (*NBEnd, int) {
	tag := NBEnd{kind: NBTypeEnd}

	fmt.Printf("-> NBEnd\n")

	return &tag, pos
}

func (c *NBEnd) GetType() byte {
	return c.kind
}

func (c *NBEnd) GetName() string {
	return ""
}
