package nbtag

import (
	"fmt"
	"log"
)

const (
	NBTypeEnd       = 0
	NBTypeByte      = 1
	NBTypeInt       = 3
	NBTypeLong      = 4
	NBTypeString    = 8
	NBTypeList      = 9
	NBTypeCompound  = 10
	NBTypeIntArray  = 11
	NBTypeLongArray = 12
)

type NBTag interface {
	GetType() byte
	GetName() string
}

type NBCompound struct {
	kind byte
	name string
}

type NBString struct {
	kind  byte
	name  string
	value string
}

type NBList struct {
	kind byte
	name string
	// TODO - what value to use here? It could be any type.
}

type NBLong struct {
	kind  byte
	name  string
	value int // TODO - does this need to be int64?
}

type NBInt struct {
	kind  byte
	name  string
	value int
}

type NBIntArray struct {
	kind  byte
	name  string
	value []int
}

type NBLongArray struct {
	kind  byte
	name  string
	value []int // TODO - does this need to be int64?
}

type NBByte struct {
	kind  byte
	name  string
	value int
}

type NBEnd struct {
	kind byte
}

func Parse(data []byte, pos int) NBTag {
	tag, _ := parseTag(data, pos)

	return tag
}

func parseTag(data []byte, pos int) (NBTag, int) {
	var tag NBTag

	t := data[pos]
	pos += 1

	switch t {
	case NBTypeEnd:
		tag, pos = parseEndTag(data, pos)
	case NBTypeByte:
		tag, pos = parseByteTag(data, pos)
	case NBTypeInt:
		tag, pos = parseIntTag(data, pos)
	case NBTypeLong:
		tag, pos = parseLongTag(data, pos)
	case NBTypeCompound:
		tag, pos = parseCompoundTag(data, pos)
	case NBTypeList:
		tag, pos = parseListTag(data, pos)
	case NBTypeString:
		tag, pos = parseStringTag(data, pos)
	case NBTypeIntArray:
		tag, pos = parseIntArrayTag(data, pos)
	case NBTypeLongArray:
		tag, pos = parseLongArrayTag(data, pos)
	default:
		log.Fatalf("-> Unhandled NBT tag type %d at pos 0x%x.\n", t, pos-1)
	}

	return tag, pos
}

func parseCompoundTag(data []byte, pos int) (*NBCompound, int) {
	tag := NBCompound{kind: NBTypeCompound}

	tag.name, pos = parseString(data, pos)

	fmt.Printf("-> NBCompound, name='%s'\n", tag.name)

	var child NBTag
	for ok := true; ok; ok = (child.GetType() != NBTypeEnd) {
		child, pos = parseTag(data, pos)
		// TODO - add child to compound's map
	}

	return &tag, pos
}

func parseEndTag(data []byte, pos int) (*NBEnd, int) {
	tag := NBEnd{kind: NBTypeEnd}

	fmt.Printf("-> NBEnd\n")

	return &tag, pos
}

func parseByteTag(data []byte, pos int) (*NBByte, int) {
	tag := NBByte{kind: NBTypeByte}

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt8(data, pos)

	fmt.Printf("-> NBByte, name='%s', value='%d'\n", tag.name, tag.value)

	return &tag, pos
}

func parseIntTag(data []byte, pos int) (*NBInt, int) {
	tag := NBInt{kind: NBTypeInt}

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt32(data, pos)

	fmt.Printf("-> NBInt, name='%s', value='%d'\n", tag.name, tag.value)

	return &tag, pos
}

func parseListTag(data []byte, pos int) (*NBList, int) {
	tag := NBList{kind: NBTypeList}

	startPos := pos

	tag.name, pos = parseString(data, pos)
	innerType, pos := parseInt8(data, pos)
	count, pos := parseInt32(data, pos)

	// TODO - for now, just skip over the bytes

	if count > 0 {
		switch innerType {
		case NBTypeList:
			for i := 0; i < count; i++ {
				for j := 1; j < 5; j++ {
					if data[pos+i+j] != 0 {
						log.Fatalf("parseListTag not yet implemented for nested lists, non-zero byte at pos=0x%x, startPos=0x%x, name='%s'", pos+i+j, startPos, tag.name)
					}
				}
			}
			pos += count * 5
		default:
			log.Fatalf("parseListTag not yet implemented for type: count=%d, innerType=%d, pos=0x%x, startPos=0x%x, name='%s'", count, innerType, pos, startPos, tag.name)
		}
	}

	fmt.Printf("-> NBList, name='%s'\n", tag.name)

	return &tag, pos
}

func parseIntArrayTag(data []byte, pos int) (*NBIntArray, int) {
	tag := NBIntArray{kind: NBTypeIntArray}

	tag.name, pos = parseString(data, pos)
	count, pos := parseInt32(data, pos)
	tag.value = make([]int, count)

	for i := 0; i < count; i++ {
		tag.value[i], pos = parseInt32(data, pos)
	}

	fmt.Printf("-> NBIntArray, name='%s', value='%d'\n", tag.name, tag.value)

	return &tag, pos
}

func parseLongArrayTag(data []byte, pos int) (*NBLongArray, int) {
	tag := NBLongArray{kind: NBTypeLongArray}

	tag.name, pos = parseString(data, pos)
	count, pos := parseInt32(data, pos)
	tag.value = make([]int, count)

	for i := 0; i < count; i++ {
		tag.value[i], pos = parseInt64(data, pos)
	}

	fmt.Printf("-> NBLongArray, name='%s', value='%d'\n", tag.name, tag.value)

	return &tag, pos
}

func parseLongTag(data []byte, pos int) (*NBLong, int) {
	tag := NBLong{kind: NBTypeLong}

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseInt64(data, pos)

	fmt.Printf("-> NBLong, name='%s', value='%d'\n", tag.name, tag.value)

	return &tag, pos
}

func parseStringTag(data []byte, pos int) (*NBString, int) {
	tag := NBString{kind: NBTypeString}

	tag.name, pos = parseString(data, pos)
	tag.value, pos = parseString(data, pos)

	fmt.Printf("-> NBString, name='%s', value='%s'\n", tag.name, tag.value)

	return &tag, pos
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

func parseInt64(data []byte, pos int) (int, int) {
	val := int(data[pos])<<56 + int(data[pos+1])<<48 + int(data[pos+2])<<40 + int(data[pos+3])<<32 + int(data[pos+4])<<24 + int(data[pos+5])<<16 + int(data[pos+6])<<8 + int(data[pos+7])
	pos += 8

	return val, pos
}

func parseInt32(data []byte, pos int) (int, int) {
	val := int(data[pos])<<24 + int(data[pos+1])<<16 + int(data[pos+2])<<8 + int(data[pos+3])
	pos += 4

	return val, pos
}

func parseInt16(data []byte, pos int) (int, int) {
	val := int(data[pos])<<8 + int(data[pos+1])
	pos += 2

	return val, pos
}

func parseInt8(data []byte, pos int) (int, int) {
	val := int(data[pos])
	pos += 1

	return val, pos
}

func (c *NBCompound) GetType() byte {
	return c.kind
}

func (c *NBCompound) GetName() string {
	return c.name
}

func (c *NBString) GetType() byte {
	return c.kind
}

func (c *NBString) GetName() string {
	return c.name
}

func (c *NBLong) GetType() byte {
	return c.kind
}

func (c *NBLong) GetName() string {
	return c.name
}

func (c *NBInt) GetType() byte {
	return c.kind
}

func (c *NBInt) GetName() string {
	return c.name
}

func (c *NBIntArray) GetType() byte {
	return c.kind
}

func (c *NBIntArray) GetName() string {
	return c.name
}

func (c *NBLongArray) GetType() byte {
	return c.kind
}

func (c *NBLongArray) GetName() string {
	return c.name
}

func (c *NBByte) GetType() byte {
	return c.kind
}

func (c *NBByte) GetName() string {
	return c.name
}

func (c *NBList) GetType() byte {
	return c.kind
}

func (c *NBList) GetName() string {
	return c.name
}

func (c *NBEnd) GetType() byte {
	return c.kind
}

func (c *NBEnd) GetName() string {
	return ""
}
