package nbtag

import "fmt"

type NBIntArray struct {
	tagData
	value []int
}

func parseIntArrayTag(data []byte, pos int) (*NBIntArray, int) {
	tag := new(NBIntArray)
	tag.kind = NBTypeIntArray

	tag.name, pos = parseString(data, pos)
	count, pos := parseInt32(data, pos)
	tag.value = make([]int, count)

	for i := 0; i < count; i++ {
		tag.value[i], pos = parseInt32(data, pos)
	}

	fmt.Printf("-> NBIntArray, name='%s', value='%d'\n", tag.name, tag.value)

	return tag, pos
}
