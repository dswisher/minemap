package nbtag

import (
	"fmt"
	"io"
	"strings"
)

type NBIntArray struct {
	tagData
	count  int
	values []int
}

func (i *NBIntArray) GetCount() int {
	return i.count
}

func (i *NBIntArray) GetValues() []int {
	return i.values
}

func newIntArrayTag() *NBIntArray {
	tag := new(NBIntArray)
	tag.kind = NBTypeIntArray

	return tag
}

// Parse an int array tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBIntArray) Parse(reader NBReader) error {
	var err error
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	return tag.parseData(reader)
}

func (tag *NBIntArray) parseData(reader NBReader) error {
	var err error

	tag.count, err = reader.ReadInt32()
	if err != nil {
		return err
	}

	tag.values = make([]int, tag.count)

	for i := 0; i < tag.count; i++ {
		v, err := reader.ReadInt32()
		if err != nil {
			return err
		}
		tag.values[i] = v
	}

	return err
}

func (tag *NBIntArray) String() string {
	return fmt.Sprintf("NBIntArray: startPos=0x%04X, count=%d, name='%s'", tag.startPos, tag.count, tag.name)
}

func (tag *NBIntArray) DumpIndented(w io.Writer, depth int) {
	const max = 8
	writeIndented(w, depth, tag.String())
	var b strings.Builder
	for i := 0; i < intMin(tag.count, max); i++ {
		if i > 0 {
			fmt.Fprintf(&b, " ")
		}
		fmt.Fprintf(&b, "%04X", tag.values[i])
	}
	if tag.count > max {
		fmt.Fprintf(&b, "...")
	}
	if b.Len() > 0 {
		writeIndented(w, depth+1, b.String())
	}
}
