package main

import (
	"bytes"
	"fmt"
	"image/gif"
)

func GifToImages(in []byte) *gif.GIF {
	reader := bytes.NewReader(in)
	images, err := gif.DecodeAll(reader)
	if err != nil {
		panic("Couldn't decode gif")
	}

	return images
}

func frameReport(g *gif.GIF) {
	for k, _ := range g.Image {
		fmt.Println(k)
	}
}

func fromFile(name string) []byte {
	var b []byte
	return b
}

func fromURL(url string) []byte {
	var b []byte
	return b
}

func main() {
	fmt.Println("Running!")
}
