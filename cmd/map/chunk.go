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

	Biomes []byte
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

	l := root.GetChild("Level")
	if l == nil {
		log.Fatalf("Chunk (%d, %d) does not contain a level!\n", cx, cz)
	}

	level, ok := l.(*nbtag.NBCompound)
	if !ok {
		log.Fatal("Level is not an NBCompound!")
	}

	// TODO - more debug fun!
	/*
		if cx == 0 && cz == 0 {
			fmt.Printf("Root, type=%d, name='%s'\n", root.GetType(), root.GetName())

			dumpTag("xPos", level.GetChild("xPos"))
			dumpTag("zPos", level.GetChild("zPos"))
			dumpTag("Biomes", level.GetChild("Biomes"))

			if level.ContainsChild("Biomes") {
				biomes := level.GetChild("Biomes").(*nbtag.NBIntArray)
				fmt.Printf("  -> Biomes, count=%d\n", biomes.GetCount())

				vals := biomes.GetValues()

				for bx := 0; bx < 16; bx++ {
					for bz := 0; bz < 16; bz++ {
						fmt.Printf("%2d ", vals[bz*16+bx])
					}
					fmt.Printf("\n")
				}
			}
		}
	*/

	// TODO - populate map data in the chunk

	// Populate some data into the chunk
	if level.ContainsChild("Biomes") {
		chunk.Biomes = make([]byte, 256)

		biomes, ok := level.GetChild("Biomes").(*nbtag.NBIntArray)
		if ok {
			vals := biomes.GetValues()
			for qq := 0; qq < 256; qq++ {
				chunk.Biomes[qq] = byte(vals[qq])
			}
		} else {
			biomes, ok := level.GetChild("Biomes").(*nbtag.NBByteArray)
			if ok {
				vals := biomes.GetValues()
				for qq := 0; qq < 256; qq++ {
					chunk.Biomes[qq] = vals[qq]
				}
			}
		}
	}

	// Return what we've built up
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
	px := c.X * 16
	pz := c.Z * 16
	var bColor color.RGBA
	if c.Biomes != nil {
		for bx := 0; bx < 16; bx++ {
			for bz := 0; bz < 16; bz++ {
				bid := c.Biomes[bz*16+bx]
				switch bid {
				case 0: // Ocean
					bColor = color.RGBA{0x00, 0x00, 0x70, 255}
				case 1: // Plains
					bColor = color.RGBA{0x8D, 0xB3, 0x60, 255}
				case 3: // Extreme Hills
					bColor = color.RGBA{0x60, 0x60, 0x60, 255}
				case 4: // Forest
					bColor = color.RGBA{0x05, 0x66, 0x21, 255}
				case 5: // Taiga
					bColor = color.RGBA{0x0B, 0x66, 0x59, 255}
				case 7: // River
					bColor = color.RGBA{0x00, 0x00, 0xFF, 255}
				case 18: // ForestHills
					bColor = color.RGBA{0x22, 0x55, 0x1C, 255}
				case 19: // TaigaHills
					bColor = color.RGBA{0x16, 0x39, 0x33, 255}
				case 34: // Wooded Mountains (1.12, Extreme Hills with Trees)
					bColor = color.RGBA{0x50, 0x70, 0x50, 255}
				default:
					log.Fatalf("Pixel color for biome id=%d not coded; chunk=(%d, %d)", bid, c.X, c.Z)
				}

				img.Set(px+bx, pz+bz, bColor)
			}
		}
	} else {
		fillSquare(img, px, pz, color.RGBA{101, 67, 33, 255}, 16)
	}
}
