package f32

import "math"

func absfloat32(x float32) float32 {
	return float32(math.Abs(float64(x)))
}

func sqrtfloat32(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}
