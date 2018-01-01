package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

type Color struct {
	RGBA uint32
}

func (c *Color) String() string {
	return fmt.Sprintf("%#x", c.RGBA)
}

func (c *Color) Set(s string) error {
	_, err := fmt.Sscanf(s, "%x", &c.RGBA)
	return err
}

func (c *Color) ToRGBA() color.RGBA {
	var mask uint32 = 0xff
	return color.RGBA{
		R: uint8((c.RGBA >> 24) & mask),
		G: uint8((c.RGBA >> 16) & mask),
		B: uint8((c.RGBA >> 8) & mask),
		A: uint8(c.RGBA & mask),
	}
}

var (
	jpgout      = flag.String("jpg", "", "output to a jpg")
	pngout      = flag.String("png", "", "output to a png")
	in          = flag.String("in", "", "input file")
	width       = flag.Int("width", 25, "width of the border")
	borderColor = new(Color)
)

func init() {
	flag.Var(borderColor, "color", "the color of the border in RGBA")
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

func applyBorder(img image.Image, c *Color, w int) image.Image {
	// Make a new solid color image, slightly larger to form the border
	r := image.Rect(0, 0, img.Bounds().Dx()+(2*w), img.Bounds().Dy()+(2*w))
	dst := image.NewRGBA(r)
	draw.Draw(dst, r, image.NewUniform(c.ToRGBA()), image.ZP, draw.Src)

	// Draw the source image over the border image
	draw.Draw(dst, r, img, image.Pt(-w, -w), draw.Src)
	return dst
}

func main() {
	flag.Parse()

	img := decode(*in)
	img = applyBorder(img, borderColor, *width)

	if *pngout != "" {
		encode(png.Encode, img, *pngout)
	}

	if *jpgout != "" {
		encode(jpegEncode, img, *jpgout)
	}
}
