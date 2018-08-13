package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"image"
	"image/color"
	"io"
	"io/ioutil"
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
		b := i * 4
		r.locations[i] = int(rawLocs[b])<<16 + int(rawLocs[b+1])<<8 + int(rawLocs[b+2])
		r.sizes[i] = rawLocs[b+3]
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
	idx := (cx & 31) + ((cz & 31) * 32)

	if r.sizes[idx] == 0 {
		return nil
	}

	// Get the position of the start of the chunk. The locations are block locations, so mult by 4096.
	pos := int64(r.locations[idx] * 0x1000)

	np, err := r.file.Seek(pos, io.SeekStart)
	if err != nil {
		log.Fatal("Error seeking to chunk location", err)
	}

	if np != pos {
		log.Fatal("Attempted to seek to %d, but got position %d.\n", pos, np)
	}

	rawHeader := make([]byte, 5)
	_, err = r.file.Read(rawHeader)
	if err != nil {
		log.Fatalf("Error reading chunk rawHeader, chunk=(%d,%d): %v", cx, cz, err)
	}

	compressedLength := int(rawHeader[0])<<24 + int(rawHeader[1])<<16 + int(rawHeader[2])<<8 + int(rawHeader[3])
	compressionType := rawHeader[4]

	if compressionType != 2 {
		log.Fatalf("Invalid compression type, chunk=(%d,%d): 0x%x", cx, cz, compressionType)
	}

	// Read the zlib compressed bytes
	compressedData := make([]byte, compressedLength)
	_, err = r.file.Read(compressedData)

	// Decompress the data
	bReader := bytes.NewReader(compressedData)
	zReader, err := zlib.NewReader(bReader)
	if err != nil {
		log.Fatalf("Error creating zlib reader, chunk=(%d,%d): %v", cx, cz, err)
	}

	defer zReader.Close()

	chunkBytes, err := ioutil.ReadAll(zReader)
	if err != nil {
		log.Fatalf("Error decompressing chunk data, chunk=(%d,%d): %v", cx, cz, err)
	}

	return ParseChunk(cx, cz, chunkBytes)
}

func (r *Region) Render(img *image.RGBA, offsetX, offsetZ int) {
	/*
		black := color.RGBA{0, 0, 0, 255}
		brown := color.RGBA{101, 67, 33, 255}
	*/

	// TODO - verify orientation - ideally, north should be towards the top of the image

	loadedChunks := 0
	nilChunks := 0

	// TODO - for the moment, only render a subsection of the region, as the parser isn't quite done
	// This works for my sample region 0,0
	for x := 0; x < 32; x++ {
		for z := 0; z < 32; z++ {
			/*
				if x >= 9 && x <= 14 && z <= 1 {
					continue
				}
				if x >= 15 && z <= 13 {
					continue
				}
			*/

			chunk := r.GetChunk(x, z)
			if chunk != nil {
				chunk.Render(img, offsetX, offsetZ)
				loadedChunks += 1
			} else {
				nilChunks += 1
			}
		}
	}

	fmt.Printf("-> Rendered region, loadedChunks=%d, nilChunks=%d\n", loadedChunks, nilChunks)
}
