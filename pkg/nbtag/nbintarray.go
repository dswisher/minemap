package nbtag

import "fmt"

type NBIntArray struct {
	tagData
	count int
	value []int
}

func (i *NBIntArray) GetCount() int {
	return i.count
}

func (i *NBIntArray) GetValues() []int {
	return i.value
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

	tag.value = make([]int, tag.count)

	for i := 0; i < tag.count; i++ {
		v, err := reader.ReadInt32()
		if err != nil {
			return err
		}
		tag.value[i] = v
	}

	return err
}

func (tag *NBIntArray) String() string {
	return fmt.Sprintf("NBIntArray: startPos=0x%04X, count=%d, name='%s'", tag.startPos, tag.count, tag.name)
}
