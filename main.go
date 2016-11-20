package main

import (
	"fmt"
	"image/jpeg"
	"os"

	"github.com/johnsudaar/sstv/encoder"
)

func main() {
	inFile, err := os.Open("/home/john/Bureau/lol.jpg")
	fmt.Println(err)
	defer inFile.Close()

	img, err := jpeg.Decode(inFile)
	fmt.Println(err)
	fmt.Println(img.Bounds())
	e := encoder.NewMartin1(96000, img)
	e.WriteHeader()
	e.EncodeImage()
	e.Sound.WriteFile("/home/john/Bureau/test.wav")
}
