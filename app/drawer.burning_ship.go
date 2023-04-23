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
	// for py := 0; py < height; py++ {
	// 	y := float64(py)/float64(height)*(ymax-ymin) + ymin
	// 	for px := 0; px < width; px++ {
	// 		x := float64(px)/float64(width)*(xmax-xmin) + xmin
	// 		z := complex(x, y)
	// 		// Image point (px, py) represents complex value z.
	// 		img.Set(px, py, (func(z complex128) color.Color {
	// 			var v complex128
	// 			for n := uint8(0); n < iterations; n++ {
	// 				v = complex(math.Abs(real(v)), math.Abs(imag(v))) + complex(math.Abs(real(z)), math.Abs(imag(z)))
	// 				if cmplx.Abs(v) > 2 {
	// 					return color.Gray{255 - contrast*n}
	// 				}
	// 			}
	// 			return color.Black
	// 		}(z)))
	// 	}
	// }

	// Dim bmp As New Bitmap(width, height)
	// For y = 0 To height - 1
	// 	For x = 0 To width - 1
	// 		Dim x0 = x_min + (x_max - x_min) * x / width
	// 		Dim y0 = y_min + (y_max - y_min) * y / height
	// 		Dim x1 = 0.0
	// 		Dim y1 = 0.0
	// 		Dim i = 0
	// 		While x1 * x1 + y1 * y1 <= 2 * 2 AndAlso i < max_iterations
	// 			Dim x2 = Math.Abs(x1 * x1 - y1 * y1 + x0)
	// 			Dim y2 = Math.Abs(2 * x1 * y1 + y0)
	// 			x1 = x2
	// 			y1 = y2
	// 			i += 1
	// 		End While
	// 		i = i * 255 / max_iterations
	// 		Dim _color = Color.FromArgb(i, i, i)
	// 		bmp.SetPixel(x, y, _color)
	// 	Next
	// Next
	// Dim file_name = burning_ship_config.output_file.Replace(".png", ".white.png")
	// bmp.Save(Path.Combine(burning_ship_config.output_directory, file_name), Imaging.ImageFormat.Png)
	// items.Add(file_name)

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

