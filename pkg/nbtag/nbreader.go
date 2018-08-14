package nbtag

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"strings"
)

type NBReader interface {
	Source() string
	Pos() int
	Context() []string

	ReadByte() (byte, error)
	ReadInt16() (int, error)
	ReadInt32() (int, error)
	ReadInt64() (int, error)
	ReadString() (string, error)
	ReadDouble() (float64, error)
	ReadFloat() (float32, error)

	PushContext(tag NBTag)
	PopContext()
}

type readerData struct {
	pos    int    // current position within the slice
	data   []byte // the underlying data
	source string // the origin of these bytes (for diagnostic purposes)

	contextStack []NBTag // Stack of tags being parsed
}

func NewReader(data []byte, source string) NBReader {
	reader := readerData{data: data, pos: 0, source: source}

	return &reader
}

func (r *readerData) ReadByte() (byte, error) {
	b := r.data[r.pos]
	r.pos += 1
	return b, nil
}

func (r *readerData) ReadInt16() (int, error) {
	val := int(binary.BigEndian.Uint16(r.data[r.pos : r.pos+2]))
	r.pos += 2
	return val, nil
}

func (r *readerData) ReadInt32() (int, error) {
	val := int(binary.BigEndian.Uint32(r.data[r.pos : r.pos+4]))
	r.pos += 4
	return val, nil
}

func (r *readerData) ReadInt64() (int, error) {
	val := int(binary.BigEndian.Uint64(r.data[r.pos : r.pos+8]))
	r.pos += 8
	return val, nil
}

func (r *readerData) ReadDouble() (float64, error) {
	bits := binary.BigEndian.Uint64(r.data[r.pos : r.pos+8])
	val := math.Float64frombits(bits)
	r.pos += 8
	return val, nil
}

func (r *readerData) ReadFloat() (float32, error) {
	bits := binary.BigEndian.Uint32(r.data[r.pos : r.pos+4])
	val := math.Float32frombits(bits)
	r.pos += 4
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

	r.pos += nameLen

	return name, nil
}

func (r *readerData) Source() string {
	return r.source
}

func (r *readerData) Pos() int {
	return r.pos
}

func (r *readerData) PushContext(tag NBTag) {
	r.contextStack = append(r.contextStack, tag)
}

func (r *readerData) PopContext() {
	if len(r.contextStack) == 0 {
		log.Fatal("Context stack underflow.")
	}

	// tag := r.contextStack[len(r.contextStack)-1]
	// fmt.Printf("-> pop tag: %s\n", tag)

	r.contextStack = r.contextStack[:len(r.contextStack)-1]
}

func (r *readerData) Context() []string {
	lines := make([]string, 0)

	for depth, tag := range r.contextStack {
		lines = append(lines, strings.Repeat(" ", depth*3)+tag.String())
	}

	const lineLen = 20

	// Add the recent bytes
	var pos int
	if len(r.contextStack) > 0 {
		pos = r.contextStack[len(r.contextStack)-1].GetStartPos() - lineLen
	} else {
		pos = r.pos - 2 // go back a little for context
	}

	var a, b strings.Builder
	for i := 0; i < 4*lineLen; i++ {
		if (i%lineLen) == 0 && i > 0 {
			lines = append(lines, b.String())
			lines = append(lines, a.String())
			a.Reset()
			b.Reset()
		}

		if (i % lineLen) == 0 {
			fmt.Fprintf(&b, "%04X:", pos+i)
			fmt.Fprint(&a, "     ")
		}

		if pos+i == r.pos {
			fmt.Fprint(&b, "[")
		} else if pos+i == r.pos+1 {
			fmt.Fprint(&b, "]")
		} else {
			fmt.Fprint(&b, " ")
		}

		val := r.data[pos+i]
		if val > 32 && val < 128 {
			fmt.Fprintf(&a, "  %c", val)
		} else {
			fmt.Fprint(&a, "  .")
		}

		fmt.Fprintf(&b, "%02X", val)
	}

	lines = append(lines, b.String())
	lines = append(lines, a.String())

	return lines
}
