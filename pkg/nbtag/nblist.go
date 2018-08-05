package nbtag

import (
	"log"
)

type NBList struct {
	tagData
	innerType byte
	count     int
	items     []NBTag
}

func parseListTag(data []byte, pos int) (*NBList, int) {
	tag := new(NBList)
	tag.startPos = pos - 1
	tag.kind = NBTypeList

	tag.name, pos = parseString(data, pos)
	it, pos := parseInt8(data, pos)
	tag.innerType = byte(it)
	tag.count, pos = parseInt32(data, pos)

	tagLog("-> NBList, start=0x%x, name='%s', innerType=%d, count=%d\n",
		tag.startPos, tag.name, tag.innerType, tag.count)

	if tag.count > 0 {
		var item NBTag

		switch tag.innerType {
		case NBTypeFloat:
			for i := 0; i < tag.count; i++ {
				item, pos = parseFloatListItem(data, pos, tag.name)
				tag.items = append(tag.items, item)
			}

		case NBTypeDouble:
			for i := 0; i < tag.count; i++ {
				item, pos = parseDoubleListItem(data, pos, tag.name)
				tag.items = append(tag.items, item)
			}

		case NBTypeList:
			// TODO - for now, just skip over the bytes
			for i := 0; i < tag.count; i++ {
				for j := 1; j < 5; j++ {
					if data[pos+i+j] != 0 {
						log.Fatalf("parseListTag not yet implemented for nested lists, non-zero byte at pos=0x%x, startPos=0x%x, name='%s'",
							pos+i+j, tag.startPos, tag.name)
					}
				}
			}
			pos += tag.count * 5

		case NBTypeCompound:
			for i := 0; i < tag.count; i++ {
				item, pos = parseCompoundListItem(data, pos, tag.name)
				tag.items = append(tag.items, item)
			}

		default:
			log.Fatalf("parseListTag not yet implemented for type: count=%d, innerType=%d, pos=0x%x, startPos=0x%x, name='%s'",
				tag.count, tag.innerType, pos, tag.startPos, tag.name)
		}
	}

	return tag, pos
}
