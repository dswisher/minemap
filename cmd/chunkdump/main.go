package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dswisher/minemap/pkg/minecraft"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s cx cz\n", os.Args[0])
		return
	}

	cx, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Unable to parse cx value '%s' as an integer: %v.\n", os.Args[1], err)
		os.Exit(1)
	}

	cz, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Unable to parse cz value '%s' as an integer: %v.\n", os.Args[2], err)
		os.Exit(1)
	}

	// Find the region and load it
	rx := cx >> 5
	rz := cz >> 5

	fmt.Printf("...loading region x=%d, z=%d...\n", rx, rz)
	// TODO - OpenRegion should return err!
	r := minecraft.OpenRegion("../../DATA/save1/region", rx, rz)
	defer r.Close()

	// Load the chunk from within the region
	// TODO - GetChunk should return err!
	chunk := r.GetChunk(cx, cz)
	if chunk == nil {
		fmt.Printf("Chunk (%d,%d) could not be opened.\n", cx, cz)
		os.Exit(2)
	}

	chunk.Dump(os.Stdout)

	fmt.Printf("len(chunk.Biomes)=%d\n", len(chunk.Biomes))

	// TODO - actually dump the data
}
