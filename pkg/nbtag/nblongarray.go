package nbtag

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

func parseLongArrayTag(data []byte, pos int) (*NBLongArray, int) {
	tag := new(NBLongArray)
	tag.startPos = pos - 1
	tag.kind = NBTypeLongArray

	tag.name, pos = parseString(data, pos)
	tag.count, pos = parseInt32(data, pos)
	tag.value = make([]int, tag.count)

	for i := 0; i < tag.count; i++ {
		tag.value[i], pos = parseInt64(data, pos)
	}

	tagLog("-> NBLongArray, name='%s', count=%d, value=%v\n", tag.name, tag.count, tag.value)

	return tag, pos
}
