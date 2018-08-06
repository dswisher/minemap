package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/dswisher/minemap/pkg/nbtag"
)

type Chunk struct {
	X, Z int
}

func ParseChunk(cx, cz int, chunkBytes []byte) *Chunk {
	// TODO - debug info
	// fmt.Printf("\n\n****** CHUNK %d, %d ******\n\n\n", cx, cz)
	// fmt.Printf("-> ParseChunk (%d, %d)\n", cx, cz)

	// TODO - HACK - save chunk to a file
	/*
		chunkFileName := fmt.Sprintf("chunk-%d-%d.dat", cx, cz)
		chunkFile, err := os.Create(chunkFileName)
		if err != nil {
			log.Fatalf("Error opening chunk dump file '%s': %v", chunkFileName, err)
		}
		chunkFile.Write(chunkBytes)
		chunkFile.Close()
		fmt.Printf("-> Wrote chunk(%d,%d) bytes to %s.\n", cx, cz, chunkFileName)
	*/

	// Create the empty chunk
	chunk := Chunk{X: cx, Z: cz}

	// First tag should be compound
	root, ok := (nbtag.Parse(chunkBytes, 0)).(*nbtag.NBCompound)
	if !ok {
		log.Fatal("Root is not NBCompound!")
	}

	if cx == 0 && cz == 0 {
		fmt.Printf("Root, type=%d, name='%s'\n", root.GetType(), root.GetName())

		l := root.GetChild("Level")
		if l == nil {
			log.Fatalf("Chunk (%d, %d) does not contain a level!\n", cx, cz)
		}

		level, ok := l.(*nbtag.NBCompound)

		if !ok {
			log.Fatal("Level is not an NBCompound!")
		}

		dumpTag("xPos", level.GetChild("xPos"))
		dumpTag("zPos", level.GetChild("zPos"))
		dumpTag("Biomes", level.GetChild("Biomes"))
	}

	// TODO - populate map data in the chunk

	return &chunk
}

func dumpTag(title string, tag nbtag.NBTag) {
	if tag == nil {
		fmt.Printf("  -> tag '%s', nil\n", title)
	} else {
		fmt.Printf("  -> tag '%s', type=%d, name='%s'\n", title, tag.GetType(), tag.GetName())
	}
}

func (c *Chunk) Render(img *image.RGBA, offsetX, offsetZ int) {
	brown := color.RGBA{101, 67, 33, 255}

	px := c.X * 16
	pz := c.Z * 16

	// TODO
	fillSquare(img, px, pz, brown, 16)
}
