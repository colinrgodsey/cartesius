package f64

import (
	"math"
	"math/rand"
)

// RandomUnitVecN provides a random n-vector with uniform distribution
// for direction and magnitude of 1.
func RandomUnitVecN(dims int, r *rand.Rand) VecC {
	const min = 0.01

	vs := make([]float64, dims)
	for {
		var dot float64
		for i := range vs {
			x := r.Float64()*2 - 1
			vs[i] = x
			dot += x * x
		}
		mag := math.Sqrt(dot)
		if mag > 1.0 || mag < min {
			continue
		}
		out := NewVecN(vs...).Div(mag)
		return &out
	}
}
