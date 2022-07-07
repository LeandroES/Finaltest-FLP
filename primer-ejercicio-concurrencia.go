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

func check3(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	imgPath := "Alfajores.jpg"
	f, err := os.Open(imgPath)
	go check3(err)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	img, _, err := image.Decode(f)

	imgPath2 := "demonio.jpg"
	f2, err := os.Open(imgPath2)
	go check3(err)
	defer func(f2 *os.File) {
		err := f2.Close()
		if err != nil {

		}
	}(f2)

	img2, _, err := image.Decode(f2)

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

	start := time.Now()
	for x := 0; x < size.X; x++ {
		go func() {
			for y := 0; y < size.Y; y++ {
				pixel := img.At(x, y)
				originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
				pixel2 := img2.At(x, y)
				originalColor2 := color.RGBAModel.Convert(pixel2).(color.RGBA)
				r := uint8(math.Sqrt((math.Pow(float64(originalColor.R), 2) + math.Pow(float64(originalColor2.R), 2)) / 2))
				g := uint8(math.Sqrt((math.Pow(float64(originalColor.G), 2) + math.Pow(float64(originalColor2.G), 2)) / 2))
				b := uint8(math.Sqrt((math.Pow(float64(originalColor.B), 2) + math.Pow(float64(originalColor2.B), 2)) / 2))
				a := uint8(math.Sqrt((math.Pow(float64(originalColor.A), 2) + math.Pow(float64(originalColor2.A), 2)) / 2))
				c := color.RGBA{
					R: r, G: g, B: b, A: a,
				}
				wImg.Set(x, y, c)
			}
		}()
	}

	elapsed := time.Since(start)
	log.Printf("Time: %s", elapsed)

	ext := filepath.Ext(imgPath)
	name := strings.TrimSuffix(filepath.Base(imgPath), ext)
	newImagePath := fmt.Sprintf("%s/%s_fusion%s", filepath.Dir(imgPath), name, ext)
	fg, err := os.Create(newImagePath)
	defer func(fg *os.File) {
		err := fg.Close()
		if err != nil {

		}
	}(fg)
	go check3(err)
	err = jpeg.Encode(fg, wImg, nil)
	go check3(err)
}
