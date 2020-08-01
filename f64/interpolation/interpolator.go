package interpolation

import (
	"runtime"
	"sync"

	"github.com/colinrgodsey/cartesius/f64"
)

// Sample2D holds a 2D interpolation sample with
// value Val at position Pos.
type Sample2D struct {
	Pos f64.Vec2
	Val float64
}

// Interpolator2D represents a specific interpolator generally
// created from a set of samples and a specific algorithm.
type Interpolator2D func(pos f64.Vec2) (float64, error)

// Multi takes a channel of positions and returns a channel of results.
// Ordering is not guaranteed, and the result channel will be closed
// when all incoming positions have been processed. Only valid positions
// will be returned on the channel.
func (interp Interpolator2D) Multi(positions <-chan f64.Vec2) <-chan Sample2D {
	var wg sync.WaitGroup
	procs := runtime.GOMAXPROCS(0)
	c := make(chan Sample2D, procs*5)

	wg.Add(procs)
	for i := 0; i < procs; i++ {
		go func() {
			for pos := range positions {
				x, err := interp(pos)
				if err == nil {
					c <- Sample2D{pos, x}
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
