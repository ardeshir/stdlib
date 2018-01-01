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
	"math"
	"os"
)

var (
	jpgout      = flag.String("jpg", "", "output to a jpg")
	pngout      = flag.String("png", "", "output to a png")
	in          = flag.String("in", "", "input file")
	filter      = flag.String("filter", "", "filter to apply")
	borderWidth = flag.Int("border", 0, "border width, used in conjunction with -color")
	borderColor = new(Color)
	blur        = flag.Int("blur", 0, "blur the image with a Gaussian blur (slow!)")
)

func init() {
	flag.Var(borderColor, "color", "the color of the border in RGBA")
}

type Gaussian struct {
	kernel  []float32
	offsets []int
}

func (gaus *Gaussian) Blur(img image.Image, x, y int) color.Color {
	colors := make([]color.Color, 0, len(gaus.kernel))
	for _, yOffset := range gaus.offsets {
		for _, xOffset := range gaus.offsets {
			colors = append(colors, img.At(x+xOffset, y+yOffset))
		}
	}
	var rsum, gsum, bsum, asum float32
	for i, c := range colors {
		rgba := color.RGBAModel.Convert(c).(color.RGBA)
		factor := gaus.kernel[i]
		rsum += factor * float32(rgba.R)
		gsum += factor * float32(rgba.G)
		bsum += factor * float32(rgba.B)
		asum += factor * float32(rgba.A)
	}
	return color.RGBA{
		R: min(255, rsum),
		G: min(255, gsum),
		B: min(255, bsum),
		A: min(255, asum),
	}
}

func normalize(kernel []float32) {
	var sum float32
	for _, f := range kernel {
		sum += f
	}
	for i := range kernel {
		kernel[i] = kernel[i] / sum
	}
}

func spread(radius int) []int {
	s := make([]int, 0, 2*radius+1)
	low, high := -radius, radius
	for i := low; i <= high; i++ {
		s = append(s, i)
	}
	return s
}

func NewGaussian(radius int) *Gaussian {
	sigmaSquared := math.Pow(float64(radius)/2, 2)
	bottom := 2 * sigmaSquared
	G := func(x, y int) float32 {
		top := -(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
		exp := math.Exp(top / bottom)
		g := 1 / (2 * math.Pi * sigmaSquared) * exp
		return float32(g)
	}

	d := radius*2 + 1
	kernel := make([]float32, 0, d*d)
	rng := spread(radius)
	for _, y := range rng {
		for _, x := range rng {
			kernel = append(kernel, G(x, y))
		}
	}
	normalize(kernel)
	return &Gaussian{kernel, rng}
}

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

type Sepia struct {
	R, G, B float32
	A       uint8
}

func min(l, r float32) uint8 {
	if r > l {
		return uint8(l)
	}
	return uint8(r)
}

func (s *Sepia) RGBA() color.RGBA {
	r := min(255, s.R*0.393+s.G*0.769+s.B*0.189)
	g := min(255, s.R*0.349+s.G*0.686+s.B*0.168)
	b := min(255, s.R*0.272+s.G*0.534+s.B*0.131)
	return color.RGBA{r, g, b, s.A}
}

func NewSepia(c color.Color) *Sepia {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return &Sepia{float32(rgba.R), float32(rgba.G), float32(rgba.B), rgba.A}
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

func doBlackAndWhite(img image.Image) image.Image {
	r := img.Bounds()
	dst := image.NewGray(r)
	draw.Draw(dst, r, img, image.ZP, draw.Src)
	return dst
}

func doSepia(img image.Image) image.Image {
	r := img.Bounds()
	dst := image.NewRGBA(r)
	w, h := r.Dx(), r.Dy()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			sepia := NewSepia(img.At(x, y)).RGBA()
			dst.Set(x, y, sepia)
		}
	}
	return dst
}

func doBorder(img image.Image, c *Color, w int) image.Image {
	r := image.Rect(0, 0, img.Bounds().Dx()+(2*w), img.Bounds().Dy()+(2*w))
	dst := image.NewRGBA(r)
	draw.Draw(dst, r, image.NewUniform(c.ToRGBA()), image.ZP, draw.Src)
	draw.Draw(dst, r, img, image.Pt(-w, -w), draw.Src)
	return dst
}

func doBlur(img image.Image, radius int) image.Image {
	g := NewGaussian(radius)
	r := img.Bounds()
	dst := image.NewRGBA(r)
	w, h := r.Dx(), r.Dy()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dst.Set(x, y, g.Blur(img, x, y))
		}
	}
	return dst
}

func main() {
	flag.Parse()

	img := decode(*in)

	switch *filter {
	case "bw":
		img = doBlackAndWhite(img)
	case "sepia":
		img = doSepia(img)
	}

	if *blur > 0 {
		img = doBlur(img, *blur)
	}

	if *borderWidth > 0 {
		img = doBorder(img, borderColor, *borderWidth)
	}

	if *pngout != "" {
		encode(png.Encode, img, *pngout)
	}

	if *jpgout != "" {
		encode(jpegEncode, img, *jpgout)
	}
}
