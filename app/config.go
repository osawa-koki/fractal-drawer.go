package app

type Config struct {
	Global     *GlobalConfig     `toml:"global"`
	Mandelbrot *MandelbrotConfig `toml:"mandelbrot"`
	Julia      *JuliaConfig      `toml:"julia"`
	Tricorn    *TricornConfig    `toml:"tricorn"`
	BurningShip *BurningShipConfig `toml:"burning_ship"`
}

type GlobalConfig struct {
	Width           int    `toml:"width"`
	Height          int    `toml:"height"`
	OutputDirectory string `toml:"output_directory"`
}

type MandelbrotConfig struct {
	Width      int     `toml:"width"`
	Height     int     `toml:"height"`
	Xmin       float64 `toml:"xmin"`
	Xmax       float64 `toml:"xmax"`
	Ymin       float64 `toml:"ymin"`
	Ymax       float64 `toml:"ymax"`
	Iterations int     `toml:"iterations"`
	Contrast   int     `toml:"contrast"`
	OutputFile string  `toml:"output_file"`
}

type JuliaConfig struct {
	Width      int     `toml:"width"`
	Height     int     `toml:"height"`
	Xmin       float64 `toml:"xmin"`
	Xmax       float64 `toml:"xmax"`
	Ymin       float64 `toml:"ymin"`
	Ymax       float64 `toml:"ymax"`
	C_real     float64 `toml:"c_real"`
	C_imag     float64 `toml:"c_imag"`
	Iterations int     `toml:"iterations"`
	Contrast   int     `toml:"contrast"`
	OutputFile string  `toml:"output_file"`
}

type TricornConfig struct {
	Width      int     `toml:"width"`
	Height     int     `toml:"height"`
	Xmin       float64 `toml:"xmin"`
	Xmax       float64 `toml:"xmax"`
	Ymin       float64 `toml:"ymin"`
	Ymax       float64 `toml:"ymax"`
	Iterations int     `toml:"iterations"`
	Contrast   int     `toml:"contrast"`
	OutputFile string  `toml:"output_file"`
}

type BurningShipConfig struct {
	Width      int     `toml:"width"`
	Height     int     `toml:"height"`
	Xmin       float64 `toml:"xmin"`
	Xmax       float64 `toml:"xmax"`
	Ymin       float64 `toml:"ymin"`
	Ymax       float64 `toml:"ymax"`
	Iterations int     `toml:"iterations"`
	Contrast   int     `toml:"contrast"`
	OutputFile string  `toml:"output_file"`
}
