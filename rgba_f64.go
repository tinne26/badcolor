package badcolor

import "math"
import "image/color"

type RGBAf64 struct {
	R float64
	G float64
	B float64
	A uint32
}

func toRGBAf64Adapter(clr color.Color) color.Color {
	return ToRGBAf64(clr)
}
var RGBAf64Model color.Model = color.ModelFunc(toRGBAf64Adapter)

func ToRGBAf64(clr color.Color) RGBAf64 {
	rgbaf64, isRGBAf64 := clr.(RGBAf64)
	if isRGBAf64 { return rgbaf64 }

	r, g, b, a := clr.RGBA()
	return RGBAf64{
		R: float64(r)/65535.0,
		G: float64(g)/65535.0,
		B: float64(b)/65535.0,
		A: a,
	}
}

func (self RGBAf64) RGBA() (r, g, b, a uint32) {
	r = uint32(math.Round(self.R*65535.0))
	g = uint32(math.Round(self.G*65535.0))
	b = uint32(math.Round(self.B*65535.0))
	return r, g, b, self.A
}

func (self RGBAf64) LinearizedRGBA() LinearizedRGBA {
	return LinearizedRGBA{
		R: LinearizeRGBChannel(self.R),
		G: LinearizeRGBChannel(self.G),
		B: LinearizeRGBChannel(self.B),
		A: self.A,
	}
}
