package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s x z png-file\n", os.Args[0])
		return
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Error parsing x", err)
	}
	z, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Error parsing z", err)
	}

	pngPath := os.Args[3]

	fmt.Printf("...loading region x=%d, z=%d...\n", x, z)
	r := OpenRegion("../../DATA/save1/region", x, z)
	defer r.Close()

	width := 512
	height := 512
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	fmt.Printf("...rendering region...\n")
	r.Render(img, 0, 0)

	fmt.Printf("...saving PNG to %s...\n", pngPath)

	// outputFile is a File type which satisfies Writer interface
	pngFile, err := os.Create(pngPath)
	if err != nil {
		log.Fatal("Error opening output PNG file", err)
	}
	defer pngFile.Close()

	png.Encode(pngFile, img) // NOTE: ignoring errors
}
