package main

import (
	"bytes"
	"fmt"
	"image/gif"
  "io/ioutil"
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

// Get a GIF from a URL
func fromURL(url string) []byte {
	var b []byte
	return b
}

func reverse(g gif.GIF) *gif.GIF {
  var out gif.GIF
  for i := len(g.Image); i > 0; i--  {
    out.Image = append(out.Image, g.Image[i-1])
    out.Delay = append(out.Delay, g.Delay[i-1])
  }

  frameReport(&out)

  return &out
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
  out := reverse(i)
  toFile("test/out.gif", out)
}
