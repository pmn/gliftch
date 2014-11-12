package glitches

import (
  "image/gif"
  "image"
  "image/color"
  "math/rand"
  )

// Applies scanlines
func apply_scanlines(destImage *image.Paletted) {
	bounds := destImage.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + 2 {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			destImage.Set(x, y, color.Black)
		}
	}
}

// Reverse a GIF. About as basic as it gets.
func Reverse(g gif.GIF) gif.GIF {
  var out gif.GIF
  for i := len(g.Image); i > 0; i--  {
    out.Image = append(out.Image, g.Image[i-1])
    out.Delay = append(out.Delay, g.Delay[i-1])
  }

  return out
}

// I dunno. Add some scan lines
func Scanlines(g gif.GIF) gif.GIF {
  var out gif.GIF
  for i := 0; i < len(g.Image); i++ {
    apply_scanlines(g.Image[i])
    out.Image = append(out.Image, g.Image[i])
    out.Delay = append(out.Delay, g.Delay[i])
  }

  return out
}

// Pick some random frames and jitter them
func Jitter(g gif.GIF) gif.GIF {
  framecount := len(g.Image)
  start := rand.Intn(framecount-5)

  var out gif.GIF
  j := 0;
  for i := 0; i < framecount; i++ {
    if (j >= 5) { j = 0 };
    out.Image = append(out.Image, g.Image[start+j])
    out.Delay = append(out.Delay, g.Delay[start+j])
    j++;
  }

  return out
}
