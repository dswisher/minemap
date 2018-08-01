package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"os"
	"path"
)

type Region struct {
	X         int      // regionX
	Z         int      // regionZ
	file      *os.File // keep around, to read chunks
	locations []int    // offsets to chunk locations
	sizes     []byte   // sizes of the chunks
}

func OpenRegion(dir string, x int, z int) *Region {
	r := Region{X: x, Z: z}

	r.Open(path.Join(dir, fmt.Sprintf("r.%d.%d.mca", x, z)))

	return &r
}

func (r *Region) Open(fpath string) {
	r.locations = make([]int, 1024)
	r.sizes = make([]byte, 1024)

	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		return
	}

	file, err := os.Open(fpath)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	r.file = file

	rawLocs := make([]byte, 1024*4)
	_, err = file.Read(rawLocs)
	if err != nil {
		log.Fatal("Error reading rawLocs", err)
	}

	for i := 0; i < 1024; i++ {
		r.locations[i] = int(rawLocs[i*4]<<16) + int(rawLocs[i*4+1]<<8) + int(rawLocs[i*4+2])
		r.sizes[i] = rawLocs[i*4+3]
	}
}

func (r *Region) Close() {
	if r.file != nil {
		r.file.Close()
	}
}

func fillSquare(img *image.RGBA, x, y int, c color.RGBA, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			img.Set(x+i, y+j, c)
		}
	}
}

func (r *Region) GetChunk(cx, cz int) *Chunk {
	idx := ((cx & 31) + (cz&31)*32)
	if r.sizes[idx] > 0 {
		pos := int64(r.locations[idx]) // TODO - mult by 4096 - loc is a block, not byte offset

		np, err := r.file.Seek(pos, io.SeekStart)
		if err != nil {
			log.Fatal("Error seeking to chunk location", err)
		}

		if np != pos {
			log.Fatal("Attempted to seek to %d, but got position %d.\n", pos, np)
		}

		return &Chunk{}
	} else {
		return nil
	}
}

func (r *Region) Render(img *image.RGBA, origX, origZ int) {
	black := color.RGBA{0, 0, 0, 255}
	brown := color.RGBA{101, 67, 33, 255}

	// TODO - verify orientation - north should be towards the top of the image

	for x := 0; x < 32; x++ {
		for z := 0; z < 32; z++ {
			// TODO - offset by origX and origZ
			px := x * 16
			pz := z * 16
			chunk := r.GetChunk(x, z)
			if chunk != nil {
				fillSquare(img, px, pz, brown, 16)
			} else {
				fillSquare(img, px, pz, black, 16)
			}
		}
	}
}
