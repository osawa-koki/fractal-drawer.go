package app

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func JuliaDrawer(config JuliaConfig) {
	width := config.Width
	height := config.Height

	xmin, xmax := config.Xmin, config.Xmax
	ymin, ymax := config.Ymin, config.Ymax
	cx, cy := config.C_real, config.C_imag

	iterations := config.Iterations
	contrast := config.Contrast

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, (func(z complex128) color.Color {
				var v complex128 = complex(cx, cy)
				for n := uint8(0); n < uint8(iterations); n++ {
					z = z*z + v
					if cmplx.Abs(z) > 2 {
						return color.Gray{255 - uint8(contrast)*n}
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
