package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

type Cropping struct {
	Width, Height uint
	X, Y          int
}

func (c *Cropping) String() string {
	return fmt.Sprintf("%dx%d%+d%+d", c.Width, c.Height, c.X, c.Y)
}

func (c *Cropping) Set(s string) error {
	_, err := fmt.Sscanf(s, "%dx%d%d%d", &c.Width, &c.Height, &c.X, &c.Y)
	return err
}

var (
	jpgout   = flag.String("jpg", "", "output to a jpg")
	pngout   = flag.String("png", "", "output to a png")
	in       = flag.String("in", "", "input file")
	cropping = new(Cropping)
)

func init() {
	flag.Var(cropping, "crop", "crop to perform, like imagemagick WxH[-+]x[-+]y")
}

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

func crop(img image.Image, c *Cropping) image.Image {
	r := image.Rect(0, 0, int(c.Width), int(c.Height))
	dst := image.NewRGBA(r)
	draw.Draw(dst, r, img, image.Pt(c.X, c.Y), draw.Src)
	return dst
}

func main() {
	flag.Parse()

	img := decode(*in)
	img = crop(img, cropping)

	if *pngout != "" {
		encode(png.Encode, img, *pngout)
	}

	if *jpgout != "" {
		encode(jpegEncode, img, *jpgout)
	}
}
