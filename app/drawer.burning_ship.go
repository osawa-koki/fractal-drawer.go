package app

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func BurningShipDrawer(config *BurningShipConfig) {
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
		for px := 0; px < width; px++ {
			x0 := xmin + (xmax-xmin)*float64(px)/float64(width)
			y0 := ymin + (ymax-ymin)*float64(py)/float64(height)
			x1 := 0.0
			y1 := 0.0
			i := 0
			for x1*x1+y1*y1 <= 2*2 && i < int(iterations) {
				x2 := math.Abs(x1*x1 - y1*y1 + x0)
				y2 := math.Abs(2*x1*y1 + y0)
				x1 = x2
				y1 = y2
				i += 1
			}
			i = i * 255 / int(iterations)
			_color := color.Gray{255 - contrast*uint8(i)}
			img.Set(px, py, _color)
		}
	}

	file, _ := os.Create(config.OutputFile)
	png.Encode(file, img)
	file.Close()
}
