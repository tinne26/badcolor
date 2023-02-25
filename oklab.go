package badcolor

import "math"
import "image/color"

type Oklab struct {
	L float64 // perceptual lightness
	A float64 // green/red-ness
	B float64 // blue/yellow-ness
	Alpha uint32
}

func toOklabAdapter(clr color.Color) color.Color {
	return ToOklab(clr)
}
var OklabModel color.Model = color.ModelFunc(toOklabAdapter)

func ToOklab(clr color.Color) Oklab {
	oklab, isOklab := clr.(Oklab)
	if isOklab { return oklab }

	linRGBA, isLinRGBA := clr.(LinearizedRGBA)
	if !isLinRGBA {
		linRGBA = ToLinearizedRGBA(clr)
	}

	x := math.Cbrt(0.4122214708*linRGBA.R + 0.5363325363*linRGBA.G + 0.0514459929*linRGBA.B)
	y := math.Cbrt(0.2119034982*linRGBA.R + 0.6806995451*linRGBA.G + 0.1073969566*linRGBA.B)
	z := math.Cbrt(0.0883024619*linRGBA.R + 0.2817188376*linRGBA.G + 0.6299787005*linRGBA.B)

	return Oklab{
		L: Clamp(0.2104542553*x + 0.7936177850*y - 0.0040720468*z, 0, 1),
		A: Clamp(1.9779984951*x - 2.4285922050*y + 0.4505937099*z, 0, 1),
		B: Clamp(0.0259040371*x + 0.7827717662*y - 0.8086757660*z, 0, 1),
		Alpha: linRGBA.A,
	}
}

func (self Oklab) LinearizedRGBA() LinearizedRGBA {
	var cube = func(x float64) float64 { return x*x*x }

	x := cube(self.L + 0.3963377774*self.A + 0.2158037573*self.B)
	y := cube(self.L - 0.1055613458*self.A - 0.0638541728*self.B)
	z := cube(self.L - 0.0894841775*self.A - 1.2914855480*self.B)

	return LinearizedRGBA{
		R: Clamp(+4.0767416621*x - 3.3077115913*y + 0.2309699292*z, 0, 1),
		G: Clamp(-1.2684380046*x + 2.6097574011*y - 0.3413193965*z, 0, 1),
		B: Clamp(-0.0041960863*x - 0.7034186147*y + 1.7076147010*z, 0, 1),
		A: self.Alpha,
	}
}

func (self Oklab) RGBA() (r, g, b, a uint32) {
	return self.LinearizedRGBA().RGBA()
}

func (self Oklab) RGBA8() color.RGBA {
	return self.LinearizedRGBA().RGBA8()
}

func (self Oklab) LCh() (l, c, h float64) {
	return LCh(self.L, self.A, self.B)
}

func (self Oklab) Equals(other Oklab) bool {
	return self.L == other.L && self.A == other.A && self.B == other.B && self.Alpha == other.Alpha
}
