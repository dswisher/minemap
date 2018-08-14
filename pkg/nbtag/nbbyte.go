package nbtag

import "fmt"

type NBByte struct {
	tagData
	value byte
}

func newByteTag() *NBByte {
	tag := new(NBByte)
	tag.kind = NBTypeByte

	return tag
}

// Parse a byte tag, including the name.
// The current position should be the byte following the tag type byte.
func (tag *NBByte) Parse(reader NBReader) error {
	var err error
	tag.name, err = reader.ReadString()
	if err != nil {
		return err
	}

	tag.value, err = reader.ReadByte()

	return err
}

func (tag *NBByte) String() string {
	return fmt.Sprintf("NBByte: startPos=0x%04X, val=%d, name='%s'", tag.startPos, tag.value, tag.name)
}
