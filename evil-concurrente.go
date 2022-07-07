package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	imgPath := "lena.jpg"
	f, err := os.Open(imgPath)
	go check(err)
	defer f.Close()
	img, _, err := image.Decode(f)
	size := img.Bounds().Size()
	wImg := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
	wg := new(sync.WaitGroup)
	start := time.Now()
	for y := 0; y < size.Y; y++ {
		wg.Add(1)
		y := y
		go func() {
			for x := 0; x < size.X; x++ {
				pixel := img.At(x, y)
				originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
				// Offset colors a little, adjust it to your taste

				r := uint8(originalColor.R)
				g := uint8(originalColor.G)
				b := uint8(originalColor.B)
				// average Alpha16 A: 65535
				//a := color.Alpha16{A: 45535}
				wImg.Set(x, y, color.CMYK{
					C: r,
					M: g,
					Y: b,
					K: 0,
				})
			}
			defer wg.Done()
		}()
	}
	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("RGB to Gray took %s", elapsed)

	ext := filepath.Ext(imgPath)
	name := strings.TrimSuffix(filepath.Base(imgPath), ext)
	newImagePath := fmt.Sprintf("%s/%s_evil%s", filepath.Dir(imgPath), name, ext)
	fg, err := os.Create(newImagePath)
	defer fg.Close()
	go check(err)
	err = jpeg.Encode(fg, wImg, nil)
	check(err)
}
