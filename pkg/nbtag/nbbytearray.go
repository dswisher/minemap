package nbtag

import "fmt"

type NBByteArray struct {
	tagData
	count int
	value []byte
}

func (i *NBByteArray) GetCount() int {
	return i.count
}

func (i *NBByteArray) GetValues() []byte {
	return i.value
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

	tag.value = make([]byte, tag.count)

	for i := 0; i < tag.count; i++ {
		v, err := reader.ReadByte()
		if err != nil {
			return err
		}
		tag.value[i] = v
	}

	return err
}

func (tag *NBByteArray) String() string {
	return fmt.Sprintf("NBByteArray: startPos=0x%04X, count=%d, name='%s'", tag.startPos, tag.count, tag.name)
}
