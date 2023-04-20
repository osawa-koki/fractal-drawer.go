package app

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func MandelbrotDrawer(config MandelbrotConfig) {
	width := config.Width
	height := config.Height

	xmin := config.Xmin
	xmax := config.Xmax
	ymin := config.Ymin
	ymax := config.Ymax

	iterations := uint8(config.Iterations)
	contrast := uint8(config.Contrast)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, (func(z complex128) color.Color {
				var v complex128
				for n := uint8(0); n < iterations; n++ {
					v = v*v + z
					if cmplx.Abs(v) > 2 {
						return color.Gray{255 - contrast*n}
					}
				}
				return color.Black
			}(z)))
		}
	}

	file, _ := os.Create(config.OutputFile)
	png.Encode(file, img)
	file.Close()
}
