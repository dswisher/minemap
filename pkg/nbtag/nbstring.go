package nbtag

import "fmt"

type NBString struct {
	tagData
	value string
}

func parseStringTag(data []byte, pos int) (*NBString, int) {
	tag := new(NBString)
	tag.kind = NBTypeString

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseString(data, pos)

	fmt.Printf("-> NBString, name='%s', value='%s'\n", tag.name, tag.value)

	return tag, pos
}

func parseString(data []byte, pos int) (string, int) {
	nameLen, pos := parseInt16(data, pos)

	var name string
	if nameLen > 0 {
		name = string(data[pos : pos+nameLen])
	} else {
		name = ""
	}

	pos += nameLen

	return name, pos
}
