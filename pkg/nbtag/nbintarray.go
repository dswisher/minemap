package nbtag

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

func parseIntArrayTag(data []byte, pos int) (*NBIntArray, int) {
	tag := new(NBIntArray)
	tag.startPos = pos - 1
	tag.kind = NBTypeIntArray

	tag.name, pos = parseString(data, pos)
	tag.count, pos = parseInt32(data, pos)
	tag.value = make([]int, tag.count)

	for i := 0; i < tag.count; i++ {
		tag.value[i], pos = parseInt32(data, pos)
	}

	tagLog("-> NBIntArray, name='%s', count=%d, value=%v\n", tag.name, tag.count, tag.value)

	return tag, pos
}
