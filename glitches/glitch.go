package glitches

import ("image/gif")

// Reverse a GIF. About as basic as it gets.
func Reverse(g gif.GIF) *gif.GIF {
  var out gif.GIF
  for i := len(g.Image); i > 0; i--  {
    out.Image = append(out.Image, g.Image[i-1])
    out.Delay = append(out.Delay, g.Delay[i-1])
  }

  return &out
}
