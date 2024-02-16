package badcolor

import "image/color"

// TODO: add version with "each" instead of slice, with quad and cubic interpolation, etc.
func OklabGradient(gradient []color.RGBA, fromColor, toColor color.Color) {
	// convert colors to oklab
	from, to := ToOklab(fromColor), ToOklab(toColor)

	// trivial cases
	if len(gradient) == 0 { return }
	if from.Equals(to) { // no change case
		rgba := from.RGBA8()
		for i := 0; i < len(gradient); i++ {
			gradient[i] = rgba
		}
		return
	}
	if len(gradient) == 2 { // trivial case
		gradient[0], gradient[1] = from.RGBA8(), to.RGBA8()
		return
	}
	
	// create general gradient
	gradient[0] = from.RGBA8()
	for i := 1; i < len(gradient) - 1; i++ {
		t := float64(i)/float64(len(gradient) - 1)
		gradient[i] = from.Interpolate(to, t).RGBA8()
	}
	gradient[len(gradient) - 1] = to.RGBA8()
}
