package nbtag

import (
	"log"
)

const (
	NBTypeCompound = 10
)

type NBTag interface {
	GetType() byte
}

type NBCompound struct {
	kind byte
	name string
}

func Parse(data []byte, pos int) NBTag {

	switch data[pos] {
	case NBTypeCompound:
		return parseCompound(data, pos+1)
	default:
		log.Fatalf("-> Unhandled NBT tag type: %d\n", data[pos])
	}

	return nil
}

func parseCompound(data []byte, pos int) *NBCompound {
	tag := NBCompound{kind: NBTypeCompound}

	tag.name, pos = parseName(data, pos)

	// TODO - parse the list of child tags

	return &tag
}

func parseName(data []byte, pos int) (string, int) {
	// TODO - add readShort?
	nameLen := int(data[pos])<<8 + int(data[pos+1])
	pos += 2

	// TODO - read the name
	pos += nameLen

	return "Hi", pos
}

func (c *NBCompound) GetType() byte {
	return c.kind
}
