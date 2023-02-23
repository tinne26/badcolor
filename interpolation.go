package badcolor

import "math"

// Linearly interpolates a and b at the given t, which must be in [0, 1].
func InterpLinear(a, b, t float64) float64 {
	return a + t*(b - a)
}

// Interpolates a and b at the given t, which must be in [0, 1], using a
// quadratic (or conic) bézier curve that passes through the control point
// ctrl.
func InterpQuad(a, ctrl, b, t float64) float64 {
	ac := InterpLinear(a, ctrl, t)
	cb := InterpLinear(ctrl, b, t)
	return InterpLinear(ac, cb, t)
}

// Interpolates a and b at the given t, which must be in [0, 1], using a
// cubic bézier curve that passes through the control points ctrl1 and ctrl2.
func InterpCubic(a, ctrl1, ctrl2, b, t float64) float64 {
	ac1  := InterpLinear(a, ctrl1, t)
	c1c2 := InterpLinear(ctrl1, ctrl2, t)
	c2b  := InterpLinear(ctrl2, b, t)
	return InterpQuad(ac1, c1c2, c2b, t)
}

func interpAlphaLinear(a, b uint32, t float64) uint32 {
	if a == b { return a }
	result := InterpLinear(float64(a), float64(b), t)
	return uint32(math.Round(result))
}

func interpAlphaQuad(a, ctrl, b uint32, t float64) uint32 {
	if a == b && a == ctrl { return a }
	result := InterpQuad(float64(a), float64(ctrl), float64(b), t)
	return uint32(math.Round(result))
}

func interpAlphaCubic(a, ctrl1, ctrl2, b uint32, t float64) uint32 {
	if a == b && a == ctrl1 && a == ctrl2 { return a }
	result := InterpCubic(float64(a), float64(ctrl1), float64(ctrl2), float64(b), t)
	return uint32(math.Round(result))
}
