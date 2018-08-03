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
	chunk := Chunk{X: cx, Z: cz}

	// First tag should be compound
	root := nbtag.Parse(chunkBytes, 0)

	if cx == 0 && cz == 0 {
		fmt.Printf("Root type = %d\n", root.GetType())
	}

	/*
		// TODO - HACK - print out a few bytes from an arbitrary chunk
		if cx == 0 && cz <= 5 && len(chunkBytes) > 10 {
			fmt.Printf("chunk(%d,%d), bytes: ", cx, cz)
			for i := 0; i < 20 && i < len(chunkBytes); i++ {
				fmt.Printf(" %02x", chunkBytes[i])
			}
			fmt.Print("\n")
		}
	*/

	// TODO - HACK - save a chunk to a file
	if cx == 16 && cz == 16 {
		chunkFileName := fmt.Sprintf("chunk-%d-%d.dat", cx, cz)
		chunkFile, err := os.Create(chunkFileName)
		if err != nil {
			log.Fatalf("Error opening chunk dump file '%s': %v", chunkFileName, err)
		}
		chunkFile.Write(chunkBytes)
		chunkFile.Close()
		fmt.Printf("-> Wrote chunk(%d,%d) bytes to %s.\n", cx, cz, chunkFileName)
	}

	// TODO - parse the bytes

	return &chunk
}
