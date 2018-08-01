package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

type Region struct {
	X       int // regionX
	Z       int // regionZ
	file    *os.File
	rawLocs []byte
}

func OpenRegion(dir string, x int, z int) *Region {
	r := Region{X: x, Z: z}

	r.Open(path.Join(dir, fmt.Sprintf("r.%d.%d.mca", x, z)))

	return &r
}

func (r *Region) Open(fpath string) {
	// TODO - if the file does not exist, we're done

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	r.file = file

	r.rawLocs = make([]byte, 32*32*4)
	_, err = file.Read(r.rawLocs)
	if err != nil {
		log.Fatal("Error reading rawLocs", err)
	}
}

func (r *Region) Close() {
	if r.file != nil {
		r.file.Close()
	}
}

func (r *Region) Print() {
	for x := 0; x < 32; x++ {
		for z := 0; z < 32; z++ {
			offset := 4 * ((x & 31) + (z&31)*32)
			if r.rawLocs[offset+3] > 0 {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}
