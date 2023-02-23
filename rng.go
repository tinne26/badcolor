package badcolor

import mrand "math/rand"
import crand "crypto/rand"

var internalRngSource *mrand.Rand
var internalRngSourceInitialized bool

func fetchRngSource() (*mrand.Rand, error) {
	if !internalRngSourceInitialized {
		var seedBytes [8]byte
		_, err := crand.Read(seedBytes[:]) // spec: on return, n == len(slice) iff err == nil
		if err != nil { return nil, err }

		var seed uint64
		for i := 0; i < 8; i++ {
			seed = (seed << 8) | uint64(seedBytes[i])
		}
		internalRngSource = mrand.New(mrand.NewSource(int64(seed)))
		internalRngSourceInitialized = true	
	}

	return internalRngSource, nil
}
