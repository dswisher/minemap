package nbtag

type NBByteArray struct {
	tagData
	count int
	value []byte
}

func parseByteArrayTag(data []byte, pos int) (*NBByteArray, int) {
	tag := new(NBByteArray)
	tag.startPos = pos - 1
	tag.kind = NBTypeByteArray

	tag.name, pos = parseString(data, pos)
	tag.count, pos = parseInt32(data, pos)
	tag.value = make([]byte, tag.count)

	var it int
	for i := 0; i < tag.count; i++ {
		it, pos = parseInt8(data, pos)
		tag.value[i] = byte(it)
	}

	tagLog("-> NBByteArray, name='%s', count=%d, value=%v\n", tag.name, tag.count, tag.value)

	return tag, pos
}
