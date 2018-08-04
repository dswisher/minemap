package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dswisher/minemap/pkg/nbtag"
)

type Chunk struct {
	X, Z int
}

func ParseChunk(cx, cz int, chunkBytes []byte) *Chunk {
	// TODO - HACK - save a chunk to a file
	if cx == 0 && cz == 0 {
		chunkFileName := fmt.Sprintf("chunk-%d-%d.dat", cx, cz)
		chunkFile, err := os.Create(chunkFileName)
		if err != nil {
			log.Fatalf("Error opening chunk dump file '%s': %v", chunkFileName, err)
		}
		chunkFile.Write(chunkBytes)
		chunkFile.Close()
		fmt.Printf("-> Wrote chunk(%d,%d) bytes to %s.\n", cx, cz, chunkFileName)
	}

	// Create the empty chunk
	chunk := Chunk{X: cx, Z: cz}

	// First tag should be compound
	root := nbtag.Parse(chunkBytes, 0)

	if cx == 0 && cz == 0 {
		fmt.Printf("Root, type=%d, name='%s'\n", root.GetType(), root.GetName())
	}

	return &chunk
}
