package badcolor

import "image/color"

// Utility alias that allows not importing image/color in some cases.
type RGBA = color.RGBA

// Creates a color with the given r, g, b values and full opacity (255).
func RGB(r, g, b uint8) color.RGBA {
	return color.RGBA{r, g, b, 255}
}

func RngRGB() (color.RGBA, error) {
	rngSource, err := fetchRngSource()
	if err != nil { return color.RGBA{}, err }

	fourBytes := rngSource.Uint32()
	return color.RGBA{
		R: uint8((fourBytes & 0xFF00_0000) >> 24),
		G: uint8((fourBytes & 0x00FF_0000) >> 16),
		B: uint8((fourBytes & 0x0000_FF00) >>  8),
		A: 255,
	}, nil
}

func RngRGBA() (color.RGBA, error) {
	rngSource, err := fetchRngSource()
	if err != nil { return color.RGBA{}, err }

	fourBytes := rngSource.Uint32()
	return color.RGBA{
		R: uint8((fourBytes & 0xFF00_0000) >> 24),
		G: uint8((fourBytes & 0x00FF_0000) >> 16),
		B: uint8((fourBytes & 0x0000_FF00) >>  8),
		A: uint8((fourBytes & 0x0000_00FF) >>  0),
	}, nil
}
