package f64

import (
	"math"
	"math/rand"
	"testing"
)

func testUnitVecN(t *testing.T, dims int) {
	const (
		num   = 100000
		eps   = 0.002
		rSeed = 8595784
	)

	sdTarget := 1.0 / math.Sqrt(float64(dims))
	r := rand.New(rand.NewSource(int64(rSeed + dims)))

	var sum float64
	var vals [num]float64
	for i := range vals {
		vec := RandomUnitVecN(dims, r)
		x := vec.At(0)
		vals[i] = x
		sum += x
	}
	mean := sum / num

	var squares float64
	for _, v := range vals {
		d := v - mean
		squares += d * d
	}
	variance := squares / num
	sd := math.Sqrt(variance)

	if math.Abs(mean) > eps {
		t.Fatalf("Mean is not within allowance: %v", mean)
	}

	if math.Abs(sd-sdTarget) > eps {
		t.Fatalf("SD is not within allowance: %v, target: %v", sd, sdTarget)
	}
}

func TestUnitVecN(t *testing.T) {
	for i := 1; i <= 5; i++ {
		testUnitVecN(t, i)
	}
}
