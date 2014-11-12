package main

import (
	"bytes"
	"fmt"
	"image/gif"
  "io/ioutil"
  "gliftch/glitches"
  "os"
)

func GifToImages(in []byte) gif.GIF {
	reader := bytes.NewReader(in)
	images, err := gif.DecodeAll(reader)
	if err != nil {
		panic("Couldn't decode gif")
	}

	return *images
}

func frameReport(g *gif.GIF) {
	for k, _ := range g.Image {
		fmt.Println(k)
	}
}

// Get a GIF from a file
func fromFile(name string) []byte {
	b, err := ioutil.ReadFile(name)
  if err != nil {
    panic(err)
  }
	return b
}

func toFile(name string, g *gif.GIF) {
  f, err := os.Create(name)
  if err != nil {
    panic(err)
  }

  defer f.Close()
  err = gif.EncodeAll(f, g)

  if err != nil {
    panic(err)
  }
}

func main() {
	fmt.Println("Running!")
  b := fromFile("test/in.gif")
  i := GifToImages(b)
  out := glitches.Reverse(i)
  toFile("test/out.gif", out)
  fmt.Println("done!")
}
