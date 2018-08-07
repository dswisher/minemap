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
