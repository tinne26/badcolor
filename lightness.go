package badcolor

import "math"
import "image/color"

// The returned value is in the [0, 1] range. Alpha is not considered.
// See also [Lightness]().
func Luminance(clr color.Color) float64 {
	linRGBA := ToLinearizedRGBA(clr)
	rc := linRGBA.R*0.2126
	gc := linRGBA.G*0.7152
	bc := linRGBA.B*0.0722
	return rc + gc + bc
}

// Returns a value in the [0, 1] range. Alpha is not considered.
// Uses [Luminance]() as a middle step during the calculations.
func Lightness(clr color.Color) float64 {
	lum := Luminance(clr)
	if lum <= 216.0/24389.0 {
		return lum*(24389.0/27.0)
	} else {
		return math.Pow(lum, 1.0/3.0)*116 - 16
	}
}
