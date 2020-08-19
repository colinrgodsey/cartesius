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
		v := sample[2]
		for i := range facets {
			Δp := pos.Sub(sample.Vec2())
			Δ := Δp.Mag()
			if Δ < ε {
				return v
			}
			w := math.Pow(Δ, -weightExp) * Δp.Dot(norms2D[i]) / Δ
			if w > facets[i].w {
				facets[i].w = w
				facets[i].sample = v
			}
		}
	}

	var w, v float64
	for _, f := range facets {
		w += f.w
		v += f.sample * f.w
	}
	return v / w
}

// MicroSphere2D is a 2D interpolator that uses microsphere
// projection to interpolate over non grid-aligned samples.
// The Z axis of the samples is the value that will be interpolated.
func MicroSphere2D(samples []Vec3) Function2D {
	return func(pos Vec2) (v float64, err error) {
		v = microSphere2D(pos, samples)
		if math.IsNaN(v) {
			v = 0
			err = ErrBadCoord
		}
		return
	}
}

//TODO: microsphere ND
