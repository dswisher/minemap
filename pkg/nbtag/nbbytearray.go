package nbtag

import (
	"fmt"
	"io"
	"strings"
)

type NBByteArray struct {
	tagData
	count  int
	values []byte
}

func (i *NBByteArray) GetCount() int {
	return i.count
}

func (i *NBByteArray) GetValues() []byte {
	return i.values
}

func newByteArrayTag() *NBByteArray {
	tag := new(NBByteArray)
	tag.kind = NBTypeByteArray

	return tag
}

// Parse a byte array tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBByteArray) Parse(reader NBReader) error {
	var err error
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	tag.count, err = reader.ReadInt32()
	if err != nil {
		return err
	}

	tag.values = make([]byte, tag.count)

	for i := 0; i < tag.count; i++ {
		v, err := reader.ReadByte()
		if err != nil {
			return err
		}
		tag.values[i] = v
	}

	return err
}

func (tag *NBByteArray) String() string {
	return fmt.Sprintf("NBByteArray: startPos=0x%04X, count=%d, name='%s'", tag.startPos, tag.count, tag.name)
}

func (tag *NBByteArray) DumpIndented(w io.Writer, depth int) {
	const max = 16
	writeIndented(w, depth, tag.String())
	var b strings.Builder
	for i := 0; i < intMin(tag.count, max); i++ {
		if i > 0 {
			fmt.Fprintf(&b, " ")
		}
		fmt.Fprintf(&b, "%02X", tag.values[i])
	}
	if tag.count > max {
		fmt.Fprintf(&b, "...")
	}
	if b.Len() > 0 {
		writeIndented(w, depth+1, b.String())
	}
}
