package badcolor

import "math"
import "image/color"

type LinearizedRGBA struct {
	R float64
	G float64
	B float64
	A uint32
}

func toLinearizedRGBAAdapter(clr color.Color) color.Color {
	return ToLinearizedRGBA(clr)
}
var LinearizedRGBAModel color.Model = color.ModelFunc(toLinearizedRGBAAdapter)

func ToLinearizedRGBA(clr color.Color) LinearizedRGBA {
	// identity case
	linRGBA, isLinRGBA := clr.(LinearizedRGBA)
	if isLinRGBA { return linRGBA }

	// rgba fast case
	rgba, isRGBA := clr.(color.RGBA)
	if isRGBA {
		return LinearizedRGBA{
			R: LinearizeRGB8Channel(rgba.R),
			G: LinearizeRGB8Channel(rgba.G),
			B: LinearizeRGB8Channel(rgba.B),
			A: uint32(rgba.A) << 8,
		}
	}

	// general case
	r, g, b, a := clr.RGBA()
	return LinearizedRGBA{
		R: LinearizeRGBChannel(float64(r)/65535.0),
		G: LinearizeRGBChannel(float64(g)/65535.0),
		B: LinearizeRGBChannel(float64(b)/65535.0),
		A: a,
	}
}

func (self LinearizedRGBA) RGBA() (r, g, b, a uint32) {
	rgba := self.RGBAf64()
	r64 := math.Round(rgba.R*65535.0)
	g64 := math.Round(rgba.G*65535.0)
	b64 := math.Round(rgba.B*65535.0)
	return uint32(r64), uint32(g64), uint32(b64), rgba.A
}

func (self LinearizedRGBA) RGBA8() color.RGBA {
	rgba := self.RGBAf64()
	r := math.Round(rgba.R*255.0)
	g := math.Round(rgba.G*255.0)
	b := math.Round(rgba.B*255.0)
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(rgba.A >> 8)}
}

func (self LinearizedRGBA) RGBAf64() RGBAf64 {
	return RGBAf64{
		R: LinearizeRGBChannelInv(self.R),
		G: LinearizeRGBChannelInv(self.G),
		B: LinearizeRGBChannelInv(self.B),
		A: self.A,
	}
}
