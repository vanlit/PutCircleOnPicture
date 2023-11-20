package PutCircleOnPicture

import (
	"image"
	"image/color"
	"image/draw"
)

func drawSquare(img draw.Image, x, y int, color_rgba color.RGBA, size int) {
	draw.Draw(
		img,
		image.Rect(x, y, x+size, y+size),
		&image.Uniform{color_rgba},
		image.Point{},
		draw.Over)
}

func DrawFilledCircle(img draw.Image, x0, y0, r int, c color.RGBA) {
	for r >= 0 {
		drawCircle_AngryWayForFilling(img, x0, y0, r, c)
		r -= 1
	}
}

func drawCircle_AngryWayForFilling(img draw.Image, x0, y0, r int, c color.RGBA) {
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)

	for x > y {
		img.Set(x0+x, y0+y, c)
		img.Set(x0+y, y0+x, c)
		img.Set(x0-y, y0+x, c)
		img.Set(x0-x, y0+y, c)
		img.Set(x0-x, y0-y, c)
		img.Set(x0-y, y0-x, c)
		img.Set(x0+y, y0-x, c)
		img.Set(x0+x, y0-y, c)

		// the rest of the Sets are the angriness
		img.Set(x0+(x-1), y0+y, c)
		img.Set(x0+(y-1), y0+x, c)
		img.Set(x0-(y-1), y0+x, c)
		img.Set(x0-(x-1), y0+y, c)
		img.Set(x0-(x-1), y0-y, c)
		img.Set(x0-(y-1), y0-x, c)
		img.Set(x0+(y-1), y0-x, c)
		img.Set(x0+(x-1), y0-y, c)

		img.Set(x0+x, y0+(y-1), c)
		img.Set(x0+y, y0+(x-1), c)
		img.Set(x0-y, y0+(x-1), c)
		img.Set(x0-x, y0+(y-1), c)
		img.Set(x0-x, y0-(y-1), c)
		img.Set(x0-y, y0-(x-1), c)
		img.Set(x0+y, y0-(x-1), c)
		img.Set(x0+x, y0-(y-1), c)

		img.Set(x0+(x-2), y0+y, c)
		img.Set(x0+(y-2), y0+x, c)
		img.Set(x0-(y-2), y0+x, c)
		img.Set(x0-(x-2), y0+y, c)
		img.Set(x0-(x-2), y0-y, c)
		img.Set(x0-(y-2), y0-x, c)
		img.Set(x0+(y-2), y0-x, c)
		img.Set(x0+(x-2), y0-y, c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
}
