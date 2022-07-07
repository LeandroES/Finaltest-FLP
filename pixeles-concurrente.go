package main

import (
	"bufio"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	// This example uses png.Decode which can only decode PNG images.
	catFile, err := os.Open("E:\\WorkSpacex\\Golang\\imagen.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer func(catFile *os.File) {
		err := catFile.Close()
		if err != nil {

		}
	}(catFile)

	// Consider using the general image.Decode as it can sniff and decode any registered image format.
	img, err := png.Decode(catFile)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("E:\\WorkSpacex\\Golang\\output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	w := bufio.NewWriter(file)

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {

		func() {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
				level := c.Y / 51 // 51 * 5 = 255
				if level == 5 {
					level--
				}
				fmt.Fprintf(w, "%v", levels[level])
			}
		}()
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
}
