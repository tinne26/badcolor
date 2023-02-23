package badcolor

import "math"
import "strconv"

// lookup table for rgba values
var linearizationTable [256]float64
func init() {
	for i := 0; i < 256; i++ {
		// normalize value
		normValue := float64(i)/255.0
		
		// convert gamma encoded RGB to linear
		var result float64
		if normValue >= 0.04045 {
			result = math.Pow((normValue + 0.055)/1.055, 2.4)
		} else {
			result = normValue/12.92
		}

		// store result as float32
		linearizationTable[i] = result
	}
}

func LinearizeRGB8Channel(value uint8) float64 {
	return linearizationTable[value]
}

// Converts a sRGB color channel value to a linearized sRGB value.
func LinearizeRGBChannel(normValue float64) float64 {
	// safety assertion
	if normValue > 1.0 || normValue < 0.0 {
		panic("normValue out of range [0, 1] (" + strconv.FormatFloat(normValue, 'f', 3, 64) + ")")
	}

	// conversion
	if normValue >= 0.04045 {
		return math.Pow(((normValue + 0.055)/1.055), 2.4)
	} else {
		return normValue/12.92
	}
}

// The inverse operation of [LinearizeRGBChannel]().
func LinearizeRGBChannelInv(normValue float64) float64 {
	// safety assertion
	if normValue > 1.0 || normValue < 0.0 {
		panic("normValue out of range [0, 1] (" + strconv.FormatFloat(normValue, 'f', 3, 64) + ")")
	}

	// conversion
	if normValue >= 0.0031308 {
		return 1.055*math.Pow(normValue, 1.0/2.4) - 0.055
	} else {
		return normValue*12.92
	}
}
