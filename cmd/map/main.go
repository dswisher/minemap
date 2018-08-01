package main

func main() {
	r := OpenRegion("../../DATA/save1/region", 0, 0)
	defer r.Close()

	r.Print()
}
