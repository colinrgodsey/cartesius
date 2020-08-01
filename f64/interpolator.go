package f64

import (
	"runtime"
	"sync"
)

// Interpolator2D is an interpolation method created using
// a set of samples and an interpolation algorithm.
type Interpolator2D func(pos Vec2) (float64, error)

// Multi takes a channel of positions and returns a channel of results.
// Ordering is not guaranteed, and the result channel will be closed
// when all incoming positions have been processed. Only valid positions
// will be returned on the channel.
func (interp Interpolator2D) Multi(positions <-chan Vec2) <-chan Vec3 {
	var wg sync.WaitGroup
	procs := runtime.GOMAXPROCS(0)
	c := make(chan Vec3, procs*5)

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
