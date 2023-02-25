package badcolor

import "math"

func Clamp(value, min, max float64) float64 {
	return math.Max(math.Min(value, max), min)
}
