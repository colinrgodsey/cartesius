package f64

import (
	"math"

	"github.com/colinrgodsey/cartesius/f64/filters"
)

// Grid2D creates a 2D grid-based interpolator using the provided filter.
// The Z axis of a sample is the value that will be interpolated.
func Grid2D(samples []Vec3, filter filters.GridFilter) (Function2D, error) {
	stride, offs, max, vals, err := makeGrid2d(samples)
	if err != nil {
		return nil, err
	}
	return func(pos Vec2) (v float64, err error) {
		var rPos [2]float64
		for i, p := range pos {
			if p < offs[i] || p > max[i] {
				err = ErrBadCoord
				return
			}
			rPos[i] = (p - offs[i]) / stride[i]
		}
		v = interp2d(vals, Vec2{rPos[0], rPos[1]}, filter)
		if math.IsNaN(v) {
			v = 0
			err = ErrBadCoord
		}
		return
	}, nil
}

/* x should already be offset and scaled */
func interp1d(values []float64, x float64, filter filters.GridFilter) float64 {
	var weights, sum float64
	low, high := filterRange(x, filter)
	for i := low; i <= high; i++ {
		if i < 0 || i >= len(values) {
			continue
		}
		w := filter.Kernel(float64(i) - x)
		sum += values[i] * w
		weights += w
	}
	return sum / weights
}

/* pos should already be offset and scaled */
func interp2d(values [][]float64, pos [2]float64, filter filters.GridFilter) float64 {
	var weights, sum float64
	low, high := filterRange(pos[1], filter)
	for i := low; i <= high; i++ {
		if i < 0 || i >= len(values) {
			continue
		}
		w := filter.Kernel(float64(i) - pos[1])
		sum += interp1d(values[i], pos[0], filter) * w
		weights += w
	}
	return sum / weights
}

func makeGrid2d(samples []Vec3) (stride, offs, max [2]float64, values [][]float64, err error) {
	for si, s := range samples {
		for i, p := range s.Vec2() {
			if p > max[i] || si == 0 {
				max[i] = p
			}
			if p < offs[i] || si == 0 {
				offs[i] = p
			}
		}
	}

	var num [2]int
	for _, s := range samples {
		for i, p := range s.Vec2() {
			if p-offs[i] == 0 {
				num[i]++
			}
		}
	}
	if num[0]*num[1] != len(samples) {
		err = ErrBadGrid
		return
	}
	for i := range stride {
		stride[i] = (max[i] - offs[i]) / float64(num[i]-1)
		max[i] += stride[i]
	}

	values = make([][]float64, num[1])
	for i := range values {
		values[i] = make([]float64, num[0])
	}
	for _, s := range samples {
		var idx [2]int
		for i, p := range s.Vec2() {
			v := (p - offs[i]) / stride[i]
			idx[i] = int(math.Round(v))
		}
		values[idx[1]][idx[0]] = s[2]
	}
	return
}

func filterRange(v float64, filter filters.GridFilter) (low, high int) {
	units := math.Ceil(filter.Size)
	low = int(math.Floor(v - units))
	high = int(math.Ceil(v + units))
	return
}
