package nbtag

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type NBReader interface {
	Source() string
	LastPos() int
	Context() []string

	ReadByte() (byte, error)
	ReadInt16() (int, error)
	ReadString() (string, error)

	// TODO - add the ability to push and pop context info (tags)
}

type readerData struct {
	pos     int    // current position within the slice
	data    []byte // the underlying data
	source  string // the origin of these bytes (for diagnostic purposes)
	lastPos int    // the position prior to the last read (for diagnostic purposes)
}

func NewReader(data []byte, source string) NBReader {
	reader := readerData{data: data, pos: 0, source: source}

	return &reader
}

func (r *readerData) ReadByte() (byte, error) {
	b := r.data[r.pos]
	r.lastPos = r.pos
	r.pos += 1
	return b, nil
}

func (r *readerData) ReadInt16() (int, error) {
	val := int(binary.BigEndian.Uint16(r.data[r.pos : r.pos+2]))
	r.lastPos = r.pos
	r.pos += 2
	return val, nil
}

func (r *readerData) ReadString() (string, error) {
	nameLen, err := r.ReadInt16()
	if err != nil {
		return "", err
	}

	var name string
	if nameLen > 0 {
		name = string(r.data[r.pos : r.pos+nameLen])
	} else {
		name = ""
	}

	// r.lastPos not updated here, as ReadInt16 already set it
	r.pos += nameLen

	return name, nil
}

func (r *readerData) Source() string {
	return r.source
}

func (r *readerData) LastPos() int {
	return r.lastPos
}

func (r *readerData) Context() []string {
	lines := make([]string, 0)

	// TODO - include context info pushed into the reader

	// Add the recent bytes
	// TODO - also include the ASCII representation, on the next line
	bytes := r.data[r.lastPos : r.lastPos+20]
	var b strings.Builder
	for i := 0; i < len(bytes); i++ {
		if i > 0 {
			fmt.Fprint(&b, " ")
		}
		fmt.Fprintf(&b, "%02X", bytes[i])
	}
	lines = append(lines, b.String())

	return lines
}
