package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func check4(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	porcentaje := 0.75

	imgPath := "Alfajores.jpg"
	f, err := os.Open(imgPath)
	check4(err)
	defer f.Close()

	img, _, err := image.Decode(f)

	imgPath2 := "demonio.jpg"
	f2, err := os.Open(imgPath2)
	check4(err)
	defer f2.Close()

	img2, _, err := image.Decode(f2)

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	start := time.Now()
	// loop though all the x
	for x := 0; x < size.X; x++ {
		// and now loop thorough all of this x's y
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			pixel2 := img2.At(x, y)
			originalColor2 := color.RGBAModel.Convert(pixel2).(color.RGBA)
			// Offset colors a little, adjust it to your taste
			r := uint8(math.Sqrt((math.Pow(porcentaje*float64(originalColor.R), 2) + math.Pow((1-porcentaje)*float64(originalColor2.R), 2)) / 2))
			g := uint8(math.Sqrt((math.Pow(porcentaje*float64(originalColor.G), 2) + math.Pow((1-porcentaje)*float64(originalColor2.G), 2)) / 2))
			b := uint8(math.Sqrt((math.Pow(porcentaje*float64(originalColor.B), 2) + math.Pow((1-porcentaje)*float64(originalColor2.B), 2)) / 2))
			a := uint8(math.Sqrt((math.Pow(porcentaje*float64(originalColor.A), 2) + math.Pow((1-porcentaje)*float64(originalColor2.A), 2)) / 2))
			// average
			c := color.RGBA{
				R: r, G: g, B: b, A: a,
			}
			wImg.Set(x, y, c)
		}
	}

	elapsed := time.Since(start)
	log.Printf("RGB to Gray took %s", elapsed)

	ext := filepath.Ext(imgPath)
	name := strings.TrimSuffix(filepath.Base(imgPath), ext)
	newImagePath := fmt.Sprintf("%s/%s_blend75%s", filepath.Dir(imgPath), name, ext)
	fg, err := os.Create(newImagePath)
	defer fg.Close()
	check4(err)
	err = jpeg.Encode(fg, wImg, nil)
	check4(err)
}
