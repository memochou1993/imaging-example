package main

import (
	"image/png"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func main() {
	input, _ := os.Open("example.png")
	defer input.Close()

	src, err := png.Decode(input)
	if err != nil {
		log.Fatal(err.Error())
	}

	output, _ := os.Create("example_resized.png")
	defer output.Close()

	dst := imaging.Resize(src, src.Bounds().Max.X/2, src.Bounds().Max.Y/2, imaging.Lanczos)

	if err = png.Encode(output, dst); err != nil {
		log.Fatal(err)
	}
}
