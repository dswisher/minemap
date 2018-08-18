package minecraft

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"os"

	"github.com/dswisher/minemap/pkg/nbtag"
)

type Chunk struct {
	X, Z int

	Biomes []byte

	RootTag *nbtag.NBCompound
}

func dumpChunk(cx, cz int, chunkBytes []byte) {
	chunkFileName := fmt.Sprintf("chunk-%d-%d.dat", cx, cz)
	chunkFile, err := os.Create(chunkFileName)
	defer chunkFile.Close()
	if err != nil {
		log.Fatalf("Error opening chunk dump file '%s': %v", chunkFileName, err)
	}
	chunkFile.Write(chunkBytes)
	fmt.Printf("-> Wrote chunk(%d,%d) bytes to %s.\n", cx, cz, chunkFileName)
}

// TODO - this should return an error instead of aborting
func ParseChunk(cx, cz int, chunkBytes []byte) *Chunk {
	// TODO - debug info
	// fmt.Printf("\n\n****** CHUNK %d, %d ******\n\n\n", cx, cz)
	// fmt.Printf("*** ParseChunk (%d, %d) ***\n", cx, cz)

	// Create the empty chunk
	chunk := Chunk{X: cx, Z: cz}

	// Parse the named-binary format.
	// TODO - this is a new version of parsing that will replace the nbTag.Parse method, used below.
	//        the new version uses a byte reader, so wrap the bytes in one...
	reader := nbtag.NewReader(chunkBytes, fmt.Sprintf("chunk (%d,%d)", cx, cz))
	topTag, err := nbtag.Parse(reader)
	if err != nil {
		// TODO - propagate the error upward; for now, go boom
		log.Print(err)
		e, ok := err.(nbtag.NBError)
		if ok {
			for _, line := range e.Context {
				log.Printf("   %s", line)
			}
		}
		// dumpChunk(cx, cz, chunkBytes)
		os.Exit(1)
	}

	if topTag == nil {
		log.Fatalf("topTag is nil!")
	}

	root, ok := topTag.(*nbtag.NBCompound)
	if !ok {
		log.Fatal("Root is not NBCompound!")
	}

	chunk.RootTag = root

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
			fmt.Printf("Root, type=%d, name='%s'\n", root.Type(), root.Name())

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
		fmt.Printf("  -> tag '%s', type=%d, name='%s'\n", title, tag.Type(), tag.Name())
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
				case 6: // Swampland
					bColor = color.RGBA{0x07, 0xF9, 0xB2, 255}
				case 7: // River
					bColor = color.RGBA{0x00, 0x00, 0xFF, 255}
				case 16: // Beach
					bColor = color.RGBA{0xFA, 0xDE, 0x55, 255}
				case 18: // ForestHills
					bColor = color.RGBA{0x22, 0x55, 0x1C, 255}
				case 19: // TaigaHills
					bColor = color.RGBA{0x16, 0x39, 0x33, 255}
				case 24: // Deep ocean
					bColor = color.RGBA{0x00, 0x00, 0x30, 255}
				case 25: // Stone beach
					bColor = color.RGBA{0xA2, 0xA2, 0x84, 255}
				case 29: // Roofed forest
					bColor = color.RGBA{0x40, 0x51, 0x1A, 255}
				case 34: // Wooded Mountains (1.12, Extreme Hills with Trees)
					bColor = color.RGBA{0x50, 0x70, 0x50, 255}
				case 129: // Sunflower plains
					bColor = color.RGBA{0xB5, 0xDB, 0x88, 255}
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

func (c *Chunk) Dump(w io.Writer) {
	if c.RootTag == nil {
		fmt.Fprintf(w, "Chunk.RootTag is nil.\n")
	} else {
		c.RootTag.DumpIndented(w, 0)
	}
}
