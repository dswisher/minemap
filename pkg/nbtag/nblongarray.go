package nbtag

import "fmt"

type NBLongArray struct {
	tagData
	count int
	value []int
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

	tag.value = make([]int, tag.count)

	for i := 0; i < tag.count; i++ {
		v, err := reader.ReadInt64()
		if err != nil {
			return err
		}
		tag.value[i] = v
	}

	return err
}

func (tag *NBLongArray) String() string {
	return fmt.Sprintf("NBLongArray: startPos=0x%04X, count=%d, name='%s'", tag.startPos, tag.count, tag.name)
}
