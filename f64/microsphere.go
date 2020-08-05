package f64

import (
	"math"
)

const (
	nFacets   = 96
	weightExp = 2
)

var norms2D [nFacets]Vec2

func init() {
	for i := range norms2D {
		Θ := float64(i) / nFacets * 2.0 * math.Pi
		norms2D[i] = Vec2{math.Cos(Θ), math.Sin(Θ)}
	}
}

func microSphere2D(pos Vec2, samples []Vec3) float64 {
	var facets [nFacets]struct {
		w, sample float64
	}
	for _, sample := range samples {
		for i := range facets {
			Δp := pos.Sub(sample.Vec2())
			Δ := Δp.Mag()
			if Δ < ε {
				return sample[2]
			}
			w := math.Pow(Δ, -weightExp) * Δp.Unit().Dot(norms2D[i])
			if w > facets[i].w {
				facets[i].w = w
				facets[i].sample = sample[2]
			}
		}
	}

	var totalW, totalV float64
	for _, f := range facets {
		totalW += f.w
		totalV += f.sample * f.w
	}
	return totalV / totalW
}

// MicroSphere2D is a 2D interpolator that using microsphere
// projection to interpolate over non grid-aligned samples.
// The Z axis of the samples is the value that will be interpolated.
func MicroSphere2D(samples []Vec3) Interpolator2D {
	return func(pos Vec2) (float64, error) {
		return microSphere2D(pos, samples), nil
	}
}

//TODO: microsphere ND
