package nbtag

import "fmt"

type NBLongArray struct {
	tagData
	value []int // TODO - does this need to be int64?
}

func parseLongArrayTag(data []byte, pos int) (*NBLongArray, int) {
	tag := new(NBLongArray)
	tag.kind = NBTypeLongArray

	tag.name, pos = parseString(data, pos)
	count, pos := parseInt32(data, pos)
	tag.value = make([]int, count)

	for i := 0; i < count; i++ {
		tag.value[i], pos = parseInt64(data, pos)
	}

	fmt.Printf("-> NBLongArray, name='%s', value='%d'\n", tag.name, tag.value)

	return tag, pos
}
