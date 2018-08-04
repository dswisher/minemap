package nbtag

import (
	"fmt"
	"log"
)

type NBList struct {
	tagData
	// TODO - what value to use here? It could be any type.
}

func parseListTag(data []byte, pos int) (*NBList, int) {
	tag := new(NBList)
	tag.kind = NBTypeList

	startPos := pos

	tag.name, pos = parseString(data, pos)
	innerType, pos := parseInt8(data, pos)
	count, pos := parseInt32(data, pos)

	if count > 0 {
		switch innerType {
		case NBTypeList:
			// TODO - for now, just skip over the bytes
			for i := 0; i < count; i++ {
				for j := 1; j < 5; j++ {
					if data[pos+i+j] != 0 {
						log.Fatalf("parseListTag not yet implemented for nested lists, non-zero byte at pos=0x%x, startPos=0x%x, name='%s'", pos+i+j, startPos, tag.name)
					}
				}
			}
			pos += count * 5
		case NBTypeCompound:
			// TODO - save the children!
			_, pos = parseCompoundData(data, pos)
		default:
			log.Fatalf("parseListTag not yet implemented for type: count=%d, innerType=%d, pos=0x%x, startPos=0x%x, name='%s'", count, innerType, pos, startPos, tag.name)
		}
	}

	fmt.Printf("-> NBList, name='%s'\n", tag.name)

	return tag, pos
}
