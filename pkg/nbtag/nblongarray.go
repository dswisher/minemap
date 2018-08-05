package nbtag

type NBLongArray struct {
	tagData
	count int
	value []int
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
