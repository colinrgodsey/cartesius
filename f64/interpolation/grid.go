package interpolation

import (
	"errors"
	"math"

	"github.com/colinrgodsey/cartesius/f64"
	"github.com/colinrgodsey/cartesius/f64/filters"
)

var (
	// ErrBadGrid is returned if the samples given don't form a grid
	ErrBadGrid = errors.New("provided samples do not form a grid")

	// ErrBadCoord is returned if you interpolate coords outside of the sample range
	ErrBadCoord = errors.New("provided coordinate is outside of the sample range")

	// ErrNotEnough is returned if less than 9 samples are provided
	ErrNotEnough = errors.New("need to provide at least 9 samples")
)

// Grid2D creates a 2D grid-based interpolator using the provided filter.
func Grid2D(samples []Sample2D, filter filters.GridFilter) Interpolator2D {
	stride, offs, max, values, err := makeGrid2d(samples)
	if err != nil {
		return func(pos f64.Vec2) (float64, error) {
			return 0, err
		}
	}
	return func(pos f64.Vec2) (float64, error) {
		var rPos [2]float64
		for i, p := range pos {
			if p < offs[i] || p > max[i] {
				return 0, ErrBadCoord
			}
			rPos[i] = (p - offs[i]) / stride[i]
		}

		v := interp2d(values, f64.Vec2{rPos[0], rPos[1]}, filter)
		return v, nil
	}
}

/* x should already be offset and scaled */
func interp1d(values []float64, x float64, filter filters.GridFilter) float64 {
	low, high := filterRange(x, filter)

	var weights, sum float64
	for i := low; i <= high; i++ {
		if i < 0 || i >= len(values) {
			continue
		}
		w := filter.Kernel(x - float64(i))
		weights += w
		sum += values[i] * w
	}
	return sum / weights
}

/* pos should already be offset and scaled */
func interp2d(values [][]float64, pos f64.Vec2, filter filters.GridFilter) float64 {
	low, high := filterRange(pos[1], filter)

	var weights, sum float64
	for i := low; i <= high; i++ {
		if i < 0 || i >= len(values) {
			continue
		}
		w := filter.Kernel(pos[1] - float64(i))
		weights += w
		sum += interp1d(values[i], pos[0], filter) * w
	}
	return sum / weights
}

func makeGrid2d(samples []Sample2D) (stride, offs, max [2]float64, values [][]float64, err error) {
	if len(samples) < 9 {
		err = ErrNotEnough
		return
	}
	for si, s := range samples {
		for i, p := range s.Pos {
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
		for i, p := range s.Pos {
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
		for i, p := range s.Pos {
			v := (p - offs[i]) / stride[i]
			idx[i] = int(math.Round(v))
		}
		values[idx[1]][idx[0]] = s.Val
	}
	return
}

func filterRange(v float64, filter filters.GridFilter) (low, high int) {
	units := math.Ceil(filter.Size)
	low = int(math.Floor(v - units))
	high = int(math.Ceil(v + units))
	return
}
