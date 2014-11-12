package main

import (
	"bytes"
	"fmt"
	"image/gif"
  "io/ioutil"
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

// Get a GIF from a file
func fromFile(name string) []byte {
	b,err := ioutil.ReadFile(name)
  if err != nil {
    panic(err)
  }
	return b
}

// Get a GIF from a URL
func fromURL(url string) []byte {
	var b []byte
	return b
}

func main() {
	fmt.Println("Running!")
  b := fromFile("test/in.gif")
  i := GifToImages(b)
  frameReport(i)
}
