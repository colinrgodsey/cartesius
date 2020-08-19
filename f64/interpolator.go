package f64

import (
	"errors"
	"runtime"
	"sync"
)

var (
	// ErrBadGrid is returned if the samples given don't form a grid
	ErrBadGrid = errors.New("provided samples do not form a grid")

	// ErrBadCoord is returned if you interpolate coords outside of the sample range
	ErrBadCoord = errors.New("provided coordinate is outside of the sample range")

	// ErrNotEnough is returned if less than 9 samples are provided
	ErrNotEnough = errors.New("need to provide at least 9 samples")
)

// Interpolator2D is an interpolation method created using
// a set of samples and an interpolation algorithm.
type Function2D func(pos Vec2) (float64, error)

// Multi takes a channel of positions and returns a channel of results.
// Ordering is not guaranteed, and the result channel will be closed
// when all incoming positions have been processed. Only valid positions
// will be returned on the channel.
func (interp Function2D) Multi(positions <-chan Vec2) <-chan Vec3 {
	var wg sync.WaitGroup
	procs := runtime.GOMAXPROCS(0)
	c := make(chan Vec3, procs*4)

	wg.Add(procs)
	for i := 0; i < procs; i++ {
		go func() {
			for pos := range positions {
				x, err := interp(pos)
				if err == nil {
					c <- Vec3{pos[0], pos[1], x}
				}
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	return c
}

// Fallback allows you to create a new interpolator that will use the
// interpolation from next if the interpolation from interp fails.
func (interp Function2D) Fallback(next Function2D) Function2D {
	return func(pos Vec2) (res float64, err error) {
		res, err = interp(pos)
		if err != nil {
			res, err = next(pos)
		}
		return
	}
}

// Grid2DPositions returns a channel of 2d positions
// across the defined grid.
func Grid2DPositions(min, stride, max Vec2) <-chan Vec2 {
	positions := make(chan Vec2, runtime.GOMAXPROCS(0)*2)
	go func() {
		for x := min[0]; x < max[0]; x += stride[0] {
			for y := min[1]; y < max[1]; y += stride[1] {
				positions <- Vec2{x, y}
			}
		}
		close(positions)
	}()
	return positions
}
