package nbtag

import (
	"fmt"
	"io"
	"strings"
)

type NBLongArray struct {
	tagData
	count  int
	values []uint
}

func newLongArrayTag() *NBLongArray {
	tag := new(NBLongArray)
	tag.kind = NBTypeLongArray

	return tag
}

// Parse a long array tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBLongArray) Parse(reader NBReader) error {
	var err error
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	tag.count, err = reader.ReadInt32()
	if err != nil {
		return err
	}

	tag.values = make([]uint, tag.count)

	for i := 0; i < tag.count; i++ {
		v, err := reader.ReadInt64()
		if err != nil {
			return err
		}
		tag.values[i] = uint(v)
	}

	return err
}

func (tag *NBLongArray) String() string {
	return fmt.Sprintf("NBLongArray: startPos=0x%04X, count=%d, name='%s'", tag.startPos, tag.count, tag.name)
}

func (tag *NBLongArray) DumpIndented(w io.Writer, depth int) {
	const max = 4
	writeIndented(w, depth, tag.String())
	var b strings.Builder
	for i := 0; i < intMin(tag.count, max); i++ {
		if i > 0 {
			fmt.Fprintf(&b, " ")
		}
		fmt.Fprintf(&b, "%08X", tag.values[i])
	}
	if tag.count > max {
		fmt.Fprintf(&b, "...")
	}
	if b.Len() > 0 {
		writeIndented(w, depth+1, b.String())
	}
}
