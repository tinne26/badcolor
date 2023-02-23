package badcolor

import "math"

// Returns the lightness, chroma and hue components of a *LAB color
// (e.g. CIELAB, Oklab, SRLAB2) by converting it to polar form. The
// hue component is returned in radians.
func LCh(l, a, b float64) (lightness, chroma, hue float64) {
	return l, math.Cbrt(a*a + b*b), math.Atan2(b, a)
}

// Converts a *LAB color (e.g. CIELAB, Oklab, SRLAB2) polar form
// components to LAB form. The hue component must be given in
// radians.
func LAB(l, c, hRads float64) (lightness, a, b float64) {
	return l, c*math.Cos(hRads), c*math.Sin(hRads)
}
