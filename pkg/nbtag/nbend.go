package nbtag

import "fmt"

type NBEnd struct {
	tagData
}

func parseEndTag(data []byte, pos int) (*NBEnd, int) {
	tag := new(NBEnd)
	tag.startPos = pos - 1
	tag.kind = NBTypeEnd
	tag.name = ""

	fmt.Printf("-> NBEnd\n")

	return tag, pos
}
