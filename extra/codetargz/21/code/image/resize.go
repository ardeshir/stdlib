package main

import (
	"flag"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

var (
	jpgout = flag.String("jpg", "", "output to a jpg")
	pngout = flag.String("png", "", "output to a png")
	in     = flag.String("in", "", "input file")
	size   = flag.Int("size", 0, "the new max dimension")
)

type encf func(io.Writer, image.Image) error

func encode(encoder encf, img image.Image, filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("failed opening %s: %s", filename, err)
		return
	}
	defer file.Close()
	err = encoder(file, img)
	if err != nil {
		log.Printf("failed encoding to %s: %s", filename, err)
	}
}

func jpegEncode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, &jpeg.Options{Quality: 80})
}

func round(value float32) int {
	if value < 0.0 {
		value -= 0.5
	} else {
		value += 0.5
	}
	return int(value)
}

func scale(w, h, size int) (int, int, float32) {
	var factor float32
	width, height := float32(w), float32(h)
	if width > height {
		factor = float32(size) / width
	} else {
		factor = float32(size) / height
	}
	return round(factor * width), round(factor * height), factor
}

func resize(img image.Image, nsize int) image.Image {
	osize := img.Bounds().Size()
	nwidth, nheight, factor := scale(osize.X, osize.Y, nsize)
	nimg := image.NewRGBA(image.Rect(0, 0, nwidth, nheight))
	for y := 0; y < nheight; y++ {
		for x := 0; x < nwidth; x++ {
			// The important stuff
			fx, fy := round(float32(x)/factor), round(float32(y)/factor)
			nimg.Set(x, y, img.At(fx, fy))
		}
	}
	return nimg
}

func decode(filename string) image.Image {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("failed decoding image: %s", err)
	}
	return img
}

func main() {
	flag.Parse()
	if *size <= 0 {
		log.Fatalln("size must be greater than 0")
	}
	img := decode(*in)
	img = resize(img, *size)

	if *pngout != "" {
		encode(png.Encode, img, *pngout)
	}

	if *jpgout != "" {
		encode(jpegEncode, img, *jpgout)
	}
}
