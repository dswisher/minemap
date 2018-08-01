package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s x z\n", os.Args[0])
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

	fmt.Printf("...loading region x=%d, z=%d...\n", x, z)
	r := OpenRegion("../../DATA/save1/region", x, z)
	defer r.Close()

	r.Print()
}
